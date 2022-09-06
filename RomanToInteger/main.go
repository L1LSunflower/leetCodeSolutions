package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
	s = "LVII"
	fmt.Println(romanToInt(s))
	s = "III"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	result := 0
	dividedOnRomanNumbers := divideOnRomanNumbers(s, []int{})
	for i := 0; i < len(dividedOnRomanNumbers); i++ {
		if i <= len(dividedOnRomanNumbers)-2 && dividedOnRomanNumbers[i+1] > dividedOnRomanNumbers[i] {
			result += dividedOnRomanNumbers[i+1] - dividedOnRomanNumbers[i]
			i++
		} else {
			result += dividedOnRomanNumbers[i]
		}
	}
	return result
}

func divideOnRomanNumbers(s string, tempSliceInts []int) []int {
	if len(s) <= 0 {
		return tempSliceInts
	}
	switch s[0] {
	case 'I':
		tempSliceInts = append(tempSliceInts, 1)
		return divideOnRomanNumbers(s[1:], tempSliceInts)
	case 'V':
		tempSliceInts = append(tempSliceInts, 5)
		return divideOnRomanNumbers(s[1:], tempSliceInts)
	case 'X':
		tempSliceInts = append(tempSliceInts, 10)
		return divideOnRomanNumbers(s[1:], tempSliceInts)
	case 'L':
		tempSliceInts = append(tempSliceInts, 50)
		return divideOnRomanNumbers(s[1:], tempSliceInts)
	case 'C':
		tempSliceInts = append(tempSliceInts, 100)
		return divideOnRomanNumbers(s[1:], tempSliceInts)
	case 'D':
		tempSliceInts = append(tempSliceInts, 500)
		return divideOnRomanNumbers(s[1:], tempSliceInts)
	case 'M':
		tempSliceInts = append(tempSliceInts, 1000)
		return divideOnRomanNumbers(s[1:], tempSliceInts)
	default:
		return tempSliceInts
	}
}
