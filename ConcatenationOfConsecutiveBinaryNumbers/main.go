package main

import "fmt"

func main() {
	num := 42
	fmt.Println(concatenatedBinary(num))
}

func concatenatedBinary(n int) int {
	var result int
	modulo := 1000000007
	for i := 1; i <= n; i++ {
		result <<= PositionsCounter(i)
		result |= i
		result %= modulo
	}
	return result
}

func PositionsCounter(n int) int {
	var counter int
	for n != 0 {
		counter++
		n >>= 1
	}
	return counter
}
