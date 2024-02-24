package main

import "fmt"

func main() {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{0, 0, 1}
	//Output: [1,3,12,0,0]
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	nonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex], nums[i] = nums[i], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}
	fmt.Println(nums)
}
