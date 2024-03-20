package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	//Output: 2, nums = [1,2,_]
	fmt.Println(removeDuplicates(nums))
}
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
		}
		nums[i] = nums[j]
	}
	for k := i + 1; k < len(nums); k++ {
		nums[k] = -1
	}
	return i + 1
}
