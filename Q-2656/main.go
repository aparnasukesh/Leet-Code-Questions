package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	//Output: 18
	fmt.Println(maximizeSum(nums, k))
}

func maximizeSum(nums []int, k int) int {

	sum := 0
	sort.Ints(nums)
	for i := 1; i <= k; i++ {
		val := nums[len(nums)-1]
		sum = sum + val
		nums[len(nums)-1] = val + 1
	}
	return sum
}
