package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	for i := 0; i <= len(nums); i++ {
		if target > nums[len(nums)/2] {
			i += len(nums) / 2
		} else {

		}
	}
	return 0
}
