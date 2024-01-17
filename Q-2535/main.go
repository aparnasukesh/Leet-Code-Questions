package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 15, 6, 3}
	//Output: 9
	fmt.Println(differenceOfSum(nums))
}

func differenceOfSum(nums []int) int {
	elemntSum := 0
	digitSum := 0
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		sum := 0
		elemntSum = elemntSum + nums[i]
		for temp > 0 {
			rem := temp % 10
			temp = temp / 10
			sum = sum + rem
		}
		digitSum = digitSum + sum
	}
	return int(math.Abs(float64(elemntSum) - float64(digitSum)))
}
