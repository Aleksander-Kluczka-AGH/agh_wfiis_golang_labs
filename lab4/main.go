package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Base struct {
	Offers
	Employees
}

type Offers struct {
	Offers []Offer `json:offers`
}

type Offer struct {
	Name      string    `json:name`
	From      time.Time `json:from`
	To        time.Time `json:to`
	Education string    `json:education`
}

type Employee struct {
	Name      string `json:name`
	Age       int    `json:age`
	Education string `json:education`
}

type Employees struct {
	People []Employee `json:people`
}

func (employees Employees) print() {
	fmt.Println("  Employees")
	for i := range employees.People {
		empl := employees.People[i]
		fmt.Println("    Employee: ", empl.Name, ", ", empl.Age, ", ", empl.Education)
	}
}

func (offers Offers) print() {
	fmt.Println("  Offers")
	for i := range offers.Offers {
		off := offers.Offers[i]
		fmt.Println("    Offer: ", off.Name, ", ", off.From, ", ", off.To, ", ", off.Education)
	}
}

type DataLoader interface {
	Load(file_name string) error
}

func (employees *Employees) Load(file_name string) error {
	file, err := os.Open(file_name)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Closing file", file_name, " failed")
		}
	}()
	if err != nil {
		fmt.Println("Opening file", file_name, " failed")
		return err
	}
	data, err := io.ReadAll(file) // data: []byte
	if err != nil {
		fmt.Println("Reading file", file_name, " failed")
		return err
	}

	// "encoding/json"
	err = json.Unmarshal(data, employees)
	if err != nil {
		fmt.Println("Deserializing contents of", file_name, " failed")
		return err
	}
	return nil
}

func (offers *Offers) Load(file_name string) error {
	file, err := os.Open(file_name)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Closing file", file_name, " failed")
		}
	}()
	if err != nil {
		fmt.Println("Opening file", file_name, " failed")
		return err
	}
	data, err := io.ReadAll(file) // data: []byte
	if err != nil {
		fmt.Println("Reading file", file_name, " failed")
		return err
	}

	// "encoding/json"
	err = json.Unmarshal(data, offers)
	if err != nil {
		fmt.Println("Deserializing contents of", file_name, " failed")
		return err
	}
	return nil
}

func (base *Base) initialize() {
	// 3
	base.Employees.Load("People.json")
	base.Offers.Load("Offers.json")
}

func (base Base) print() {
	// 5
	fmt.Println("Database:")
	base.Offers.print()
	base.Employees.print()
}

func (base *Base) readFromCLI() {
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose what to add:")
	fmt.Println("1. Person")
	fmt.Println("2. Offer")
	in.Scan()
	option := in.Text()
	switch option {
	case "1":
		{
			employee := Employee{}
			employee.readFromCLI(in)
			base.Employees.People = append(base.Employees.People, employee)
		}
	case "2":
		{
			offer := Offer{}
			offer.readFromCLI(in)
			base.Offers.Offers = append(base.Offers.Offers, offer)
		}
	}
}

func (employee *Employee) readFromCLI(buf *bufio.Scanner) {
	fmt.Println("Reading person data:")
	fmt.Print("Name = ")
	buf.Scan()
	name := buf.Text()

	fmt.Print("Age = ")
	buf.Scan()
	age, _ := strconv.Atoi(buf.Text())

	fmt.Print("Education = ")
	buf.Scan()
	edu := buf.Text()

	employee.Name = name
	employee.Age = age
	employee.Education = edu
}

func (offer *Offer) readFromCLI(buf *bufio.Scanner) {
	// 4
	fmt.Println("Reading offer data:")
	fmt.Print("Name = ")
	buf.Scan()
	name := buf.Text()

	fmt.Print("From = ")
	buf.Scan()
	from, _ := time.Parse("YYYY-MM-DD", buf.Text())

	fmt.Print("To = ")
	buf.Scan()
	to, _ := time.Parse("YYYY-MM-DD", buf.Text())

	fmt.Print("Education = ")
	buf.Scan()
	edu := buf.Text()

	offer.Name = name
	offer.From = from
	offer.To = to
	offer.Education = edu
}

func (base Base) matchingOffers() map[Employee][]Offer {
	// 6 (zwraca na ekran)
	result := map[Employee][]Offer{}
	fmt.Println("Printing matching offers to all people:")
	for i := range base.Employees.People {
		person := base.Employees.People[i]
		result[person] = make([]Offer, 0)

		fmt.Println("  Person: ", person.Name, ", ", person.Age, ", ", person.Education)
		for j := range base.Offers.Offers {
			offer := base.Offers.Offers[j]

			if offer.Education == person.Education {
				result[person] = append(result[person], offer)
				fmt.Println("    Matching offer: ", offer.Name, ", ", offer.From, ", ", offer.To, ", ", offer.Education)
			}
		}
	}
	return result
}

func (base Base) costOfEmploymentOfAllMatchedOffers() int {
	// 7
	matched := base.matchingOffers()
	result := 0
	for person, offers := range matched {
		if len(offers) != 0 {
			result += person.Age // im starszy tym drozszy
		}
	}
	fmt.Println("Total cost of employment: ", result)
	return result
}

func main() {
	base := Base{}
	base.initialize()
	base.print()
	base.readFromCLI()
	base.print()
	base.matchingOffers()
	base.costOfEmploymentOfAllMatchedOffers()
}
