package main

import "fmt"

func main() {
	var num uint32 = 00000000000000000000000000001011
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

//func hammingWeight(num uint32) int {
//	var bits = []bool{}
//	for num >= 1 {
//		fmt.Println(num % 8)
//		if num%8 != 0 {
//			bits = append(bits, true)
//		}
//		num /= 8
//	}
//	return len(bits)
//}
