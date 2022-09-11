package main

import "fmt"

func main() {
	num1, num2 := "12341341234", "532323452"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1, num2 string) string {
	var overflow bool
	var tempRes int
	if len(num1) >= len(num2) {
		numRunes := []rune(num1)
		for i := 0; i <= len(num2)-1; i++ {
			if overflow {
				tempRes = convertByteToInt(numRunes[len(numRunes)-1-i]) + convertByteToInt(rune(num2[len(num2)-1-i])) + 1
				overflow = false
			} else {
				tempRes = convertByteToInt(numRunes[len(numRunes)-1-i]) + convertByteToInt(rune(num2[len(num2)-1-i]))
			}
			if tempRes >= 10 {
				overflow = true
				tempRes -= 10
			}
			numRunes[len(numRunes)-1-i] = rune(tempRes + '0')
		}
		//fmt.Println(numRunes)
		num1 = string(numRunes)
	}
	return num1
}

func convertByteToInt(r rune) int {
	return int(r) - '0'
}
