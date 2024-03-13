package main

import "fmt"

func main() {
	nums := []int{2, 11, 10, 1, 3}
	k := 10
	//Output: 3

	fmt.Println(minOperations(nums, k))
}

func minOperations(nums []int, k int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			count++
		}
	}
	return count
}
