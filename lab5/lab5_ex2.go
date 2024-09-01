package main

import (
	"flag"
	"math/rand"
)

type Limit struct {
	limit int
}

func ex2() {
	lower_limit := Limit{0}
	upper_limit := Limit{100}
	var low_flag = flag.IntVar(&lower_limit.limit, "low", lower_limit.limit, "Lower bound of randomization range.")
	var high_flag = flag.IntVar(&upper_limit.limit, "low", upper_limit.limit, "Upper bound of randomization range.")
	var print_hist_flag = flag.Bool("hist", false, "Print histogram of randomized values to stdout.")
	flag.Parse()

	genNumber := func() int {
		return rand.Intn(*low_flag-*high_flag) + *low_flag
	}

	numbers := make([]int, 0)
	sum := 0
	for range 50 {
		number := genNumber()
		sum += number
		numbers = append(numbers, number)
	}

	avg := float32(sum) / float32(len(numbers))

	section := upper_limit.limit - lower_limit.limit
	hist_map := map[int][]int{}
	for i := range numbers {
		num := numbers[i]

	}
}

func main() {
	ex2()
}
