package main

import "fmt"

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	//"0P"
	s = "aa"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	var tempString string
	for _, r := range s {
		validDigit, rightRune := isValidDigit(r)
		if validDigit {
			tempString += string(rightRune)
		}
	}
	secondPart := ""
	for i := len(tempString) - 1; i > (len(tempString)-1)/2; i-- {
		secondPart += string(tempString[i])
	}
	//fmt.Println(tempString[:len(tempString)/2], secondPart)
	if tempString[:len(tempString)/2] == secondPart {
		return true
	} else {
		return false
	}

	//for i := 0; i < len(tempString)/2; i++ {
	//	if tempString[i] == tempString[len(tempString)-1-i] {
	//		continue
	//	} else {
	//		return false
	//	}
	//}

	//return true
}

func isValidDigit(r rune) (bool, rune) {
	if r >= 'A' && r <= 'Z' {
		return true, r + 32
	} else if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
		return true, r
	}
	return false, rune(1)
}
