package main

import "fmt"

func main() {
	start, goal := 3, 5
	fmt.Println(minBitFlips(start, goal))
}

func minBitFlips(start int, goal int) int {
	counter, tempVar := 0, 1
	for start != goal && tempVar != 0 {
		start ^= tempVar
		counter++
		if tempVar <= goal && counter <= goal/2 {
			tempVar <<= 1
		} else {
			tempVar >>= 1
		}
		fmt.Println(tempVar, start)
	}
	//fmt.Println(start)
	return counter
}
