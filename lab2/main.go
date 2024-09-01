package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	const numOfStrings = 50

	// generation of strings
	strings := []string{}
	for range numOfStrings {
		oldCap := cap(strings)

		stringLength := rand.Intn(10) + 4
		strings = append(strings, generateString(stringLength))

		if oldCap != cap(strings) {
			fmt.Println("Detected capacity change (len/cap): ", len(strings), cap(strings))
		}
	}

	// elimination of strings where 2 neighbor characters are the same
	strings = filterStrings(strings, hasNoNeighborCharacters)

	// split strings into map based on their length
	mapper := map[int][]string{}
	for i := range 14 {
		mapper[i] = make([]string, 0)
	}
	for i := range strings {
		current := strings[i]
		length := len(current)
		mapper[length] = append(mapper[length], current)
	}

	for k, v := range mapper {
		fmt.Println("mapper[", k, "] = ", v)
	}

	for k := range mapper {
		fmt.Println("mapper[", k, "] = ", mapper[k])
	}
}

func ex2() {
	var tab_length int = rand.Intn(4) + 1
	tab2d := make([]([]int), tab_length)
	for i := 0; i < tab_length; i++ {
		random_2dim_length := rand.Intn(4) + 1
		tab2d[i] = make([]int, random_2dim_length)
		for j := range random_2dim_length {
			tab2d[i][j] = rand.Intn(5)
		}
	}

	// print rows of the table
	for i := 0; i < tab_length; i++ {
		fmt.Println(tab2d[i])
	}
	fmt.Println()

	// create a set with all values from the table
	set := make(map[int]bool, 0)
	for searched := range 5 {
		all_rows_have_searched := true
		for i := 0; i < tab_length; i++ {
			row_length := len(tab2d[i])
			var row_has_searched bool = false
			for j := range row_length {
				if tab2d[i][j] == searched {
					row_has_searched = true
					break
				}
			}
			if !row_has_searched {
				all_rows_have_searched = false
				break
			}
		}
		if all_rows_have_searched {
			set[searched] = true
		}
	}

	for i := range 5 {
		_, ok := set[i]
		fmt.Println("set[", i, "] = ", ok)
	}
}

func ex3() {
	// Proszę zaimplementować funkcję wstawiającą wartość do tablicy zachowując porządek sortowania. Funkcja powinna zwracać slice oraz error (różny od nil w przypadku przekroczenia zakresu)
	tab := [5]int{1, 2, 4, 5}
	fmt.Println(tab, len(tab))

	inserted_tab, err := insert_sorted(tab[:], 3)
	fmt.Println(inserted_tab, err)
}

func ex4() {
	var str StringApplier = ident
	var target string = "hello"
	var prefix_decor StringApplier = PrefixDecorator(str, "yes, ")
	var suffix_decor StringApplier = SuffixDecorator(str, " world")
	var aggregated StringApplier = SuffixDecorator(PrefixDecorator(str, "yes, "), " world")
	fmt.Println(prefix_decor(target))
	fmt.Println(suffix_decor(target))
	fmt.Println(aggregated(target))
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

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func insert_sorted(target []int, element int) ([]int, error) {
	target_length := len(target)
	non_zero_length := 0
	for i := range target {
		if target[i] != 0 {
			non_zero_length++
		}
	}

	lower := make([]int, 0)
	greater := make([]int, 0)
	for i := range target {
		if target[i] == 0 {
			continue
		}
		if target[i] <= element {
			lower = append(lower, target[i])
		}
		if target[i] > element {
			greater = append(greater, target[i])
		}
	}
	result := append([]int{}, lower...)
	result = append(result, element)
	result = append(result, greater...)

	if len(result) > non_zero_length && target_length == non_zero_length {
		return result, errors.New("Resulting slice is greater than passed array.")
	} else {
		return result, nil
	}
}

type StringApplier func(string) string

func PrefixDecorator(fn StringApplier, prefix string) StringApplier {
	return func(s string) string {
		return fn(prefix + s)
	}
}

func SuffixDecorator(fn StringApplier, suffix string) StringApplier {
	return func(s string) string {
		return fn(s + suffix)
	}
}

func ident(s string) string {
	return s
}
