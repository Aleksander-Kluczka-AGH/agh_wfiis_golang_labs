package main

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

func a_default(input string) []string {
	return strings.Fields(input)
}

func a_regex(input string, regex string) []string {
	reg := regexp.MustCompile(regex)
	return strings.FieldsFunc(input, func(r rune) bool {
		return reg.Match([]byte{byte(r)})
	})
}

func A(input ...string) ([]string, error) {
	if len(input) < 1 || len(input) > 2 {
		return []string{}, errors.New("invalid number of arguments")
	} else if len(input) == 1 {
		return a_default(input[0]), nil
	} else {
		return a_regex(input[0], input[1]), nil
	}
}

func B(input string) ([]string, error) {
	result, err := A(input)
	sort.Strings(result)
	return result, err
}

func C(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}
	return count
}

func D(input string) int {
	words, _ := A(input)
	count := 0
	for i := range words {
		if len(words[i]) > 0 {
			count++
		}
	}
	return count
}

func E(input string) int {
	count := 0
	for i := range input {
		if input[i] != '\n' && input[i] != ' ' && input[i] != '\t' {
			count++
		}
	}
	return count
}

func F(input string) map[string]int {
	words, _ := A(input)
	word_counts := make(map[string]int)
	for i := range words {
		word_counts[words[i]] = 0
	}

	for i := range words {
		word_counts[words[i]]++
	}
	return word_counts
}

func isPalindrome(str string) bool {
	lastIdx := len(str) - 1
	for i := 0; i < (lastIdx - i); i++ {
		if str[i] != str[lastIdx-i] {
			return false
		}
	}
	return true
}

func G(input string) []string {
	words, _ := A(input)
	palindromes := make([]string, 0)

	for i := range words {
		if isPalindrome(words[i]) {
			palindromes = append(palindromes, words[i])
		}
	}
	sort.Strings(palindromes)
	return palindromes
}
