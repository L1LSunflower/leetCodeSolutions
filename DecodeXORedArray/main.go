package main

import "fmt"

func main() {
	encodedSlice := []int{6, 2, 7, 3}
	first := 4
	fmt.Println(decode(encodedSlice, first))
}

func decode(encoded []int, first int) []int {
	encoded = append([]int{first}, encoded...)
	for i := 1; i < len(encoded); i++ {
		encoded[i] = encoded[i] ^ encoded[i-1]
	}
	return encoded
}
