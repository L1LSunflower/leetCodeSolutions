package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println(toLowerCase(s))
	s = "LOVELY"
	fmt.Println(toLowerCase(s))
}

func toLowerCase(s string) string {
	result := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return result
}
