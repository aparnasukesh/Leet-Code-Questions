package main

import "fmt"

func main() {
	nums := []int{5, 6, 2, 7, 4}
	fmt.Println(maxProductDifference(nums))
	//34
}

func maxProductDifference(nums []int) int {
	var w, x, y, z int
	// for i := 0; i < len(nums); i++ {
	// 	if x < nums[i] {
	// 		x = nums[i]
	// 	}
	// }
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	w = nums[len(nums)-1]
	x = nums[len(nums)-2]
	y = nums[0]
	z = nums[1]

	return (w * x) - (y * z)
}
