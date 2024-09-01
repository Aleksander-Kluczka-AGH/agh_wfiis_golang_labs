package main

func fib_chan() <-chan int {
	ch := make(chan int)
	// iterative fibonacci
	go func() {
		a, b := 0, 1
		for {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

func even_range(min int, max int, ch <-chan int) int {
	sum := 0
	for {
		el := <-ch
		// println(el)
		if el > max {
			break
		}
		if el >= min && el <= max && (el%2 == 0) {
			sum += el
		}
	}
	return sum
}

func odd_range(min int, max int, ch <-chan int) int {
	sum := 0
	for {
		el := <-ch
		// println(el)
		if el > max {
			break
		}
		if el >= min && el <= max && (el%2 != 0) {
			sum += el
		}
	}
	return sum
}

func main() {
	fib_ch := fib_chan()
	println(even_range(0, 100, fib_ch))

	fib_ch2 := fib_chan()
	println(odd_range(0, 100, fib_ch2))

}
