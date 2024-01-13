package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 4, 2, 3}
	//Output: [3,2,5,4]
	fmt.Println(numberGame(nums))
}
func numberGame(nums []int) []int {
	game := make([]int, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			game[i+1] = nums[i]
		} else {
			game[i-1] = nums[i]
		}
	}
	return game
}
