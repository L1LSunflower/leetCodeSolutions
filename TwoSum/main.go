package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 9, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	return findSum(nums, target, 0)
}

func findSum(nums []int, target, index int) []int {
	if len(nums)-1 == index {
		return nums
	}
	for i := index + 1; i <= len(nums)-1; i++ {
		if nums[index]+nums[i] == target {
			return []int{index, i}
		}
	}
	return findSum(nums, target, index+1)
}
