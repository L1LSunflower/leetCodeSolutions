package main

import "fmt"

func main() {
	fmt.Println(isPalidrome(-121))
}

func isPalidrome(x int) bool {
	if x < 0 {
		return false
	}
	reversedNumber := divideOnDigits(x, []int{})
	for i := 0; i < len(reversedNumber); i++ {
		if reversedNumber[i] != reversedNumber[len(reversedNumber)-1-i] {
			return false
		}
	}
	return true
}

func divideOnDigits(num int, tempRes []int) []int {
	if num == 0 {
		return tempRes
	}
	tempRes = append(tempRes, num%10)
	return divideOnDigits(num/10, tempRes)
}
