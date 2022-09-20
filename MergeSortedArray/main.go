package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	nums1 = append(nums1[:m], nums2[:n]...)
//	for i := 0; i < len(nums1); i++ {
//		for j := 0; j < len(nums1)-i-1; j++ {
//			if nums1[j] > nums1[j+1] {
//				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
//			}
//		}
//	}
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1, nums2 = nums1[:m], nums2[:n]
	//if m == 0 {
	//	nums1 = nums2
	//} else {
	for i, j := 0, 0; i < len(nums1)-1; i++ {
		fmt.Println(j)
		if nums1[i] <= nums2[j] && (nums1[i+1] >= nums2[j] || nums1[i+1] == 0) {
			//fmt.Println(nums1, j)
			nums1 = append(nums1[:i+1], nums1[i:]...)
			nums1[i+1] = nums2[j]
			j++
		} else if i == len(nums1)-j {
			nums1 = append(nums1, nums2[j:]...)
		}
	}
	//}
}
