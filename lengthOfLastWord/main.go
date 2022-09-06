package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
	s = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	var count int

	for i := len(s) - 1; 0 <= i; i-- {
		if validDigit(s[i]) {
			count++
		} else if count > 0 && !validDigit(s[i]) {
			return count
		} else {
			continue
		}
	}

	return count
}

func validDigit(b byte) bool {
	if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' {
		return true
	}
	return false
}
