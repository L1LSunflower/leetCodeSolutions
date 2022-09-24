package main

import "fmt"

func main() {
	n := 123
	fmt.Println(numberOfSteps(n))
}

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		if num&1 == 1 && num != 1 {
			count += 2
		} else {
			count++
		}
		num >>= 1
	}
	return count
}
