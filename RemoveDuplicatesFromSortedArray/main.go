package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

//func removeDuplicates(nums []int) int {
//
//	return len(nums)
//}

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			//fmt.Println(nums[:i], nums[i+1:])
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	//fmt.Println(nums)
	return len(nums)
}
