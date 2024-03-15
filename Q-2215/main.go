package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	//Output: [[3],[]]
	fmt.Println(findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	a, b := uniqueMap(nums1), uniqueMap(nums2)
	return [][]int{
		isNotIn(a, b),
		isNotIn(b, a),
	}
}

func isNotIn(a map[int]bool, b map[int]bool) []int {
	out := make([]int, 0)
	for num := range a {
		if _, v2 := b[num]; !v2 {
			out = append(out, num)
		}
	}
	return out
}

func uniqueMap(nums []int) map[int]bool {
	unique := make(map[int]bool)
	for _, num := range nums {
		if _, v := unique[num]; !v {
			unique[num] = true
		}
	}
	return unique
}
