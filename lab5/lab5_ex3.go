package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func initLogger() {
	log.SetPrefix("LOG:")
	//log.SetFlags(log.Ldate|log.Lmicroseconds|log.Longfile)
	//log.Lshortfile
	log.SetFlags(log.LstdFlags)
}

type Songs struct {
	Songs []Song `json:songs`
}

type Song struct {
	Rank   int    `json:rank`
	Title  string `json:title`
	Artist string `json:artist`
	Album  string `json:album`
	Year   string `json:year`
}

func (songs *Songs) readJsonFile(file_name string) error {
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

	// "encoding/json"
	err = json.Unmarshal(data, songs)
	if err != nil {
		log.Println("Deserializing contents of", file_name, " failed")
		return err
	}
	return nil
}

func ex3() {
	initLogger()
	var sort_by_flag = flag.String("sort_key", "", "Key name which the contents will be sorted by.")
	var count_flag = flag.Int("count", 5, "Number of records to display (default: 5).")

	flag.Parse()

	if *sort_by_flag == "" {
		log.Fatalln("Sorting key -sort_key must be specified.")
	}
	log.Println("count flag value =", *count_flag)

	songs := Songs{}
	songs.readJsonFile("songs.json")

	sort.Slice(songs.Songs, func(i, j int) bool {
		switch strings.ToLower(*sort_by_flag) {
		case "rank":
			return songs.Songs[i].Rank < songs.Songs[j].Rank
		case "title":
			return songs.Songs[i].Title < songs.Songs[j].Title
		case "artist":
			return songs.Songs[i].Artist < songs.Songs[j].Artist
		case "album":
			return songs.Songs[i].Album < songs.Songs[j].Album
		case "year":
			return songs.Songs[i].Year < songs.Songs[j].Year
		default:
			return false
		}
	})

	fmt.Println(songs.Songs[0:*count_flag])
}

func main() {
	ex3()
}
