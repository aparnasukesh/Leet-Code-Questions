package main

import "fmt"

func main() {
	nums := []int{0, 1, 4, 6, 7, 10}
	diff := 3
	//Output: 2
	fmt.Println(arithmeticTriplets(nums, diff))
}

func arithmeticTriplets(nums []int, diff int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := i + 2; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					count++
				}
			}
		}
	}
	return count
}
