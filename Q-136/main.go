package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	singleOne := singleNumber(nums)
	fmt.Println("Single Number is :", singleOne)
}

func singleNumber(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		count = 1
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] == nums[j] {
				count = count + 1
			}
		}
		if count == 1 {
			return nums[i]
		}
	}
	return 0
}
