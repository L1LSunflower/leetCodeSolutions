package main

import "fmt"

func main() {
	operations := []string{"--X", "X++", "X++"}
	fmt.Println(finalValueAfterOperations(operations, 0))
	//operations := []string{"--X", "X++", "X++"}
	//fmt.Println(finalValueAfterOperations(operations))
	operations = []string{"X++", "++X", "--X", "X--"}
	fmt.Println(finalValueAfterOperations(operations, 0))
}

func finalValueAfterOperations(operations []string, result int) int {
	if len(operations) <= 0 {
		return result
	} else if operations[0] == "X++" || operations[0] == "++X" {
		return finalValueAfterOperations(operations[1:], result+1)
	} else {
		return finalValueAfterOperations(operations[1:], result-1)
	}
}

//func finalValueAfterOperations(operations []string) int {
//	result := 0
//	for _, v := range operations {
//		if v == "--X" || v == "X--" {
//			result--
//		} else {
//			result++
//		}
//	}
//	return result
//}
