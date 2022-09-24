package main

import "fmt"

func main() {
	n, start := 5, 0
	fmt.Println(xorOperation(n, start))
}

func xorOperation(n int, start int) int {
	counter := 1
	result := start
	for counter < n {
		result ^= start + 2*counter
		counter++
	}
	return result
}
