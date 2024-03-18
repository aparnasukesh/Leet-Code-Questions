package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	//Output: true
	fmt.Println(containsNearbyDuplicate(nums, k))
}
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, val := range nums {
		index, ok := m[val]
		if ok && (i-index) <= k {
			return true
		}
		m[val] = i
	}
	return false
}
