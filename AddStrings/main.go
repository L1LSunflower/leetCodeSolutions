package main

import "fmt"

func main() {
	num1, num2 := "1", "9"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1, num2 string) string {
	if len(num1) > len(num2) {
		return sumStringValues(num1, num2)
	} else {
		return sumStringValues(num2, num1)
	}
}

func sumStringValues(num1, num2 string) string {
	var overflow bool
	var tempRes int
	numRunes := []rune(num1)
	for i := 0; i < len(num2) || overflow == true; i++ {
		//fmt.Println(numRunes)
		if overflow {
			if i >= len(num2) && i < len(numRunes) {
				tempRes = convertByteToInt(numRunes[len(numRunes)-1-i]) + 1
			} else if i >= len(numRunes) {
				return "1" + string(numRunes)
			} else {
				tempRes = convertByteToInt(numRunes[len(numRunes)-1-i]) + convertByteToInt(rune(num2[len(num2)-1-i])) + 1
			}
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
	num1 = string(numRunes)

	return num1
}

func convertByteToInt(r rune) int {
	return int(r) - '0'
}
