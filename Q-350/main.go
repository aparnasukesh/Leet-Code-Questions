package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}

	// nums1 := []int{3, 1, 3}
	// nums2 := []int{1, 1}

	res := intersect(nums1, nums2)
	fmt.Println(res)
}

func intersect(nums1 []int, nums2 []int) []int {

	m1 := make(map[int]int)

	for _, val := range nums1 {
		m1[val]++
	}
	fmt.Println(m1)

	var ans []int
	for _, v := range nums2 {
		if m1[v] > 0 {
			ans = append(ans, v)
			m1[v]--
		}
	}
	return ans
}
