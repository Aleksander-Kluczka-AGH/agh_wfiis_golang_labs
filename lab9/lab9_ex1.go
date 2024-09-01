package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func spam(delay time.Duration) {
	chars := `-\|/`
	i := 0
	for {
		fmt.Print("\r", string(chars[i]))
		time.Sleep(delay)
		i = (i + 1) % len(chars)
	}

}

func main() {
	go spam(100 * time.Millisecond)
	fmt.Println("\r", fib(45))
}
