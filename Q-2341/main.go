package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 1, 3, 2, 2}
	//Output: [3,1]
	fmt.Println(numberOfPairs(nums))
}
func numberOfPairs(nums []int) []int {
	res := make([]int, 2)
	count := 0
	for i := 0; i < len(nums); i++ {
		num1 := 0
		num2 := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				num1 = i
				num2 = j
				count++
			}
		}
		fmt.Println(num1, num2)
		fmt.Println(nums)

		fmt.Println(nums)
		fmt.Println(nums)
		res[0] = count
	}
	res[1] = len(nums)
	fmt.Println(res)
	return res
}
