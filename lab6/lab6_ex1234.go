package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"

	"golang.org/x/term"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []*User  `xml:"user"`
}

type User struct {
	XMLName  xml.Name `xml:"user"`
	Login    string   `xml:"login"`
	Password string   `xml:"paswword"`
	Role     int      `xml:"role"`
}

type Person struct {
	XMLName    xml.Name `xml:"person"`
	Id         int      `xml:"id"`
	FirstName  string   `xml:"firstName"`
	LastName   string   `xml:"lastName"`
	Age        int      `xml:"age"`
	Birth      Data     `xml:"birth"`
	Death      Data     `xml:"death"`
	Pesel      string   `xml:"pesel"`
	CreditCard string   `xml:"creditcard"`
	Gender     string   `xml:"gender"`
}

type Data struct {
	D, M, Y int
}

type People struct {
	XMLName xml.Name  `xml:"persons"`
	People  []*Person `xml:"person"`
}

type Role int

const (
	NONE Role = 0
	READ Role = 1
	EDIT Role = 2
	ADD  Role = 4
	ALL  Role = 7
)

var users Users
var people People
var systemUsersXmlFile *os.File
var peopleXmlFile *os.File
var is_logged_in bool = false
var logged_in_user_role Role = NONE

func (users *Users) readXmlFile(file_name string) error {
	file, err := os.Open(file_name)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalln("Closing file", file_name, " failed")
		}
	}()
	if err != nil {
		log.Fatalln("Opening file", file_name, " failed")
		return err
	}
	data, err := io.ReadAll(file) // data: []byte
	if err != nil {
		log.Fatalln("Reading file", file_name, " failed")
		return err
	}

	err = xml.Unmarshal(data, users)
	if err != nil {
		log.Println("Deserializing contents of", file_name, " failed")
		return err
	}
	return nil
}

func (people *People) readFromBytestring(data []byte) {
	err := xml.Unmarshal(data, people)
	if err != nil {
		log.Println("Deserializing contents of bytestring failed")
	}
}

func (users *Users) writeXmlFile(file *os.File) {
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")
	err := encoder.Encode(users)
	if err != nil {
		fmt.Println("Error encoding XML:", err)
	}
}

func (users Users) print() {
	for _, user := range users.Users {
		fmt.Println("Login: ", user.Login)
	}
}

func createUsersXml() {
	xmlFile, err := os.Create("users.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer xmlFile.Close()
}

func readPasswordFromCli() string {
	fmt.Println("Enter password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
	}
	sha256Hash := sha256.Sum256(password)
	passh := hex.EncodeToString(sha256Hash[:])
	return passh
}

func registerNewSystemUser(users *Users, file *os.File) {
	var login string
	var role int
	fmt.Println("Enter login: ")
	fmt.Scan(&login)
	var password string = readPasswordFromCli()
	fmt.Println("Enter role: ")
	fmt.Println("0 - NONE")
	fmt.Println("1 - READ")
	fmt.Println("2 - EDIT")
	fmt.Println("4 - ADD")
	fmt.Println("7 - ALL")
	fmt.Scan(&role)

	user := User{Login: login, Password: password, Role: role}
	users.Users = append(users.Users, &user)

	// save users to users.xml
	users.writeXmlFile(file)
}

func safeOpenFile(file_name string) *os.File {
	file, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalln("Opening file", file_name, " failed")
	}
	return file
}

func readEncryptedFile(file_name string) []byte {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalln("Opening file", file_name, " failed")
	}
	data, err := io.ReadAll(file) // data: []byte
	if err != nil {
		log.Fatalln("Reading file", file_name, " failed")
	}
	peopleXmlFile = file

	return data
}

func decryptBytestring(cipherText []byte, key string) []byte {
	var key_byte []byte = []byte("The giraffes enter the wardrobe.")

	// decrypt file
	block, err := aes.NewCipher(key_byte)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	return plainText
}

func readDatabase() {
	cipherText := readEncryptedFile("encrypted.xml")
	decoded := decryptBytestring(cipherText, "The giraffes enter the wardrobe.")
	people.readFromBytestring(decoded)
}

func menu(users *Users, file *os.File) bool {
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Actions:")
	fmt.Println("1. Register new system user")
	fmt.Println("2. Login as system user")
	fmt.Println("3. Exit")
	in.Scan()
	option := in.Text()
	switch option {
	case "1":
		{
			registerNewSystemUser(users, file)
		}
	case "2":
		{
			login()
		}
	case "3":
		{
			return false
		}
	}
	return true
}

func login() {
	fmt.Println("Enter login:")
	users.print()

	var login string
	fmt.Scan(&login)

	var password string = readPasswordFromCli()

	for _, user := range users.Users {
		if user.Login == login && user.Password == password {
			is_logged_in = true
			logged_in_user_role = Role(user.Role)
			return
		}
	}

	fmt.Println("Login failed: invalid login or password")
}

func initApp() {
	if _, err := os.Stat("users.xml"); errors.Is(err, os.ErrNotExist) {
		log.Println("users.xml does not exist. Creating a new one")
		createUsersXml()
		systemUsersXmlFile = safeOpenFile("users.xml")
		users.writeXmlFile(systemUsersXmlFile)

	} else {
		log.Println("users.xml exists. Reading from it")
		xmlFile := safeOpenFile("users.xml")
		users.readXmlFile("users.xml")
		systemUsersXmlFile = xmlFile
	}

	readDatabase()
}

func deinit() {
	var err error
	err = systemUsersXmlFile.Close()
	if err != nil {
		log.Fatalln("Closing file users.xml failed")
	}
	err = peopleXmlFile.Close()
	if err != nil {
		log.Fatalln("Closing file encrypted.xml failed")
	}
}

func main() {
	initApp()
	for {
		if !menu(&users, systemUsersXmlFile) {
			break
		}
	}
	deinit()
}
