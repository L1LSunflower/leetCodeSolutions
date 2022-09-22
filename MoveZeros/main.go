package main

import "fmt"

func main() {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
	// [0,1,0,3,12]
	nums = []int{0, 1, 0, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			count++
		}
		fmt.Println(count)
	}
	for i := 0; i < count; i++ {
		nums = append(nums, 0)
	}
}
