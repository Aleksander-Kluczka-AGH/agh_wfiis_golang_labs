package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func initLogger() {
	log.SetPrefix("LOG:")
	//log.SetFlags(log.Ldate|log.Lmicroseconds|log.Longfile)
	//log.Lshortfile
	log.SetFlags(log.LstdFlags)
}

func readAndParse(input io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func uniqueWords(content []string) []string {
	log.Println("Selecting unique words")
	unique_map := make(map[string]bool)

	for _, sentence := range content {
		words := strings.Fields(sentence)
		for _, word := range words {
			unique_map[word] = true
		}
	}

	unique_words := make([]string, 0, len(unique_map))
	for word := range unique_map {
		unique_words = append(unique_words, word)
	}

	return unique_words
}

func filteredWords(content []string, prefix string) []string {
	log.Println("Filtering words")
	var filteredWords []string

	for _, sentence := range content {
		sentenceWords := strings.Fields(sentence)
		for _, word := range sentenceWords {
			if !strings.HasPrefix(word, prefix) {
				filteredWords = append(filteredWords, word)
			}
		}
	}

	return filteredWords
}

func ex1() {
	initLogger()
	log.Println("Starting zad1")

	var op_flag = flag.String("op", "", "Choose operation to perform on input. Options: 'uniq', 'filter'.")
	var stdin_src_flag = flag.Bool("stdin", false, "Read from stdin instead of predefined file.")
	var file_flag = flag.String("file", "", "Path to the file to be read when -stdin switch has not been specified.")

	flag.Parse()

	var content io.Reader

	if *stdin_src_flag == true {
		// from stdin
		content = os.Stdin
	} else {
		// from file
		if len(*file_flag) == 0 {
			log.Fatalln("When reading from file, source must be specified.")
		}
		file, _ := os.Open(*file_flag)
		content = file
		defer func() {
			err := file.Close()
			if err != nil {
				fmt.Println("Closing file example_file.txt failed")
			}
		}()
	}

	lines, _ := readAndParse(content)

	if *op_flag == "uniq" {
		unique_content := uniqueWords(lines)
		fmt.Println(unique_content)
	} else if *op_flag == "filter" {
		filtered_content := filteredWords(lines, "n")
		fmt.Println(filtered_content)
	} else {
		log.Fatalln("Flag '-op' is not specified")
	}
}

func main() {
	ex1()
}
