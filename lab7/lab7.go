package main

import (
	"fmt"
	"strconv"
)

func one[T any](args []T, pred func(T) bool) []T {
	var result []T
	for _, arg := range args {
		if pred(arg) {
			result = append(result, arg)
		}
	}
	return result
}

func two[T any](args []T, pred func(T) T) []T {
	var result []T
	for _, arg := range args {
		result = append(result, pred(arg))
	}
	return result
}

func three[T any](args []T, pred func(T, T) T) T {
	var result T = pred(args[0], args[1])
	for i, arg := range args {
		if i >= 2 {
			result = pred(result, arg)
		}
	}
	return result
}

func four[T map[K]V, K comparable, V any](arg T) ([]K, []V) {
	var keys []K
	var values []V
	for key, val := range arg {
		keys = append(keys, key)
		values = append(values, val)
	}
	return keys, values
}

type FivePair[K any, V any] struct {
	key   K
	value V
}

func five[T map[K]V, K comparable, V any](arg T) []FivePair[K, V] {
	var result []FivePair[K, V]
	for key, val := range arg {
		result = append(result, FivePair[K, V]{key: key, value: val})
	}
	return result
}

func six[KS []K, VS []V, K comparable, V any](keys KS, values VS) map[K]V {
	var result map[K]V = make(map[K]V)
	var length int = min(len(keys), len(values))
	for iter := range length {
		result[keys[iter]] = values[iter]
	}
	return result
}

func seven[T int | float64](args []string) ([]T, error) {
	convertedSlice := make([]interface{}, len(args))
	var t interface{} = *new(T)
	switch t.(type) {
	case int:
		for i, arg := range args {
			if val, err := strconv.Atoi(arg); err != nil {
				return nil, err
			} else {
				convertedSlice[i] = val
			}
		}
	case float64:
		for i, arg := range args {
			if val, err := strconv.ParseFloat(arg, 64); err != nil {
				return nil, err
			} else {
				convertedSlice[i] = val
			}
		}
	}

	resultSlice := make([]T, len(convertedSlice))
	for i, val := range convertedSlice {
		resultSlice[i] = val.(T)
	}
	return resultSlice, nil
}

func main() {
	// one
	fmt.Println(one([]int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 }))

	// two
	fmt.Println(two([]int{1, 2, 3, 4, 5}, func(x int) int { return x * x }))

	// three
	fmt.Println(three([]int{1, 2, 3, 4, 5}, func(x int, y int) int { return x + y }))

	// four
	fmt.Println(four(map[int]float32{1: 1.1, 2: 2.2, 3: 3.3}))

	// five
	fmt.Println(five(map[int]float32{1: 1.1, 2: 2.2, 3: 3.3}))

	// six
	fmt.Println(six([]int{1, 2, 3, 4, 5}, []float32{1.1, 2.2, 3.3}))

	// seven
	fmt.Println(seven[int]([]string{"1", "2", "3", "4", "5"}))
}
