package main

import "fmt"

func main() {
	num := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(num, target)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {

	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
			}
		}
	}
	return res
}
