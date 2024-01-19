package main

import "fmt"

func main() {
	nums1 := []int{4, 3, 2, 3, 1}
	nums2 := []int{2, 2, 5, 2, 3, 6}
	//Output: [3,4]

	fmt.Println(findIntersectionValues(nums1, nums2))
}
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	res := make([]int, 2)
	arr1 := make([]int, len(nums1))
	arr2 := make([]int, len(nums2))
	count1 := 0
	count2 := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if arr2[j] != 1 {
				if nums1[i] == nums2[j] {
					count1++
					arr2[j] = 1
				}
			}

		}
	}
	for i := 0; i < len(nums2); i++ {
		for j := 0; j < len(nums1); j++ {
			if arr1[j] != 1 {
				if nums2[i] == nums1[j] {
					count2++
					arr1[j] = 1
				}
			}
		}
	}
	res[0] = count2
	res[1] = count1
	return res
}
