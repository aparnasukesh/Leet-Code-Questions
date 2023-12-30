package main

import "fmt"

func main() {
	nums := []int{-100, -98, -1, 2, 3, 4}
	fmt.Println(maximumProduct(nums))
}

func maximumProduct(nums []int) int {
	product := 1
	one := 0
	// two := 0
	// three := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] <= nums[j] {
				one = nums[j]
			}
		}
	}
	fmt.Println(one)
	return product
}
