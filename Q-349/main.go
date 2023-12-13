package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
func intersection(nums1 []int, nums2 []int) []int {
	var length1, length2 int
	res := []int{}
	var m, n []int
	if len(nums1) <= len(nums2) {
		length1 = len(nums1)
		length2 = len(nums2)
		m = nums1
		n = nums2

	} else {
		length1 = len(nums2)
		length2 = len(nums1)
		m = nums2
		n = nums1
	}
	check := make(map[int]bool, length1)

	for i := 0; i < length1; i++ {
		for j := 0; j < length2; j++ {
			if m[i] == n[j] {
				if _, ok := check[m[i]]; !ok {
					check[m[i]] = true
					res = append(res, m[i])
				}
			}

		}
	}

	return res
}
