package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 4, 5, 2}
	//Output: 12
	fmt.Println(maxProduct(nums))
}
func maxProduct(nums []int) int {
	var product int
	length := len(nums)
	if len(nums) >= 3 {
		sort.Ints(nums)
		product = (nums[length-1] - 1) * (nums[length-2] - 1)
	} else {
		product = (nums[length-1] - 1) * (nums[length-2] - 1)
	}

	return product

}
