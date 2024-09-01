package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	ex1()
	ex2()
}

func ex1() {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(reflect.TypeOf(letters).Kind())
	fmt.Println(reflect.TypeOf(letters[0]).Kind())
}

func ex2() {
	const numOfStrings = 50

	// generation of strings
	strings := []string{}
	for range numOfStrings {
		oldCap := cap(strings)

		stringLength := rand.Intn(10) + 3
		strings = append(strings, generateString(stringLength))

		if oldCap != cap(strings) {
			fmt.Println("Detected capacity change (len/cap): ", len(strings), cap(strings))
		}
	}

	// elimination of strings where 2 neighbor characters are the same
	strings = filterStrings(strings, hasNoNeighborCharacters)

	// split strings into 3 slices based on their length
	short := []string{}
	medium := []string{}
	long := []string{}
	for i := range strings {
		current := strings[i]
		length := len(current)
		switch { // empty switch
		case length < 7:
			short = append(short, current)
		case 7 <= length && length < 11:
			medium = append(medium, current)
		case 11 <= length && length < 14:
			long = append(long, current)
		}

		switch length { // switch by length
		case 3, 4, 5, 6:
			// short = append(short, current)
		case 7, 8, 9, 10:
			// medium = append(medium, current)
		case 11, 12, 13:
			// long = append(long, current)
		}
	}

	fmt.Println("short strings = ", short)
	fmt.Println("medium strings = ", medium)
	fmt.Println("long strings = ", long)
}

func generateString(length int) string {
	const letters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func hasNoNeighborCharacters(str string) bool {
	var iter int = 0
	for iter < (len(str) - 1) {
		if str[iter] == str[iter+1] {
			return false
		}
		iter++
	}
	return true
}

type Predicate[A any] func(A) bool

func filterStrings(strings []string, pred Predicate[string]) []string {
	result := []string{}
	for i := range strings {
		if pred(strings[i]) {
			result = append(result, strings[i])
		}
	}
	return result
}
