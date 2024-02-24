package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{10, 4, 8, 3}
	//Output: [15,1,11,22]
	fmt.Println(leftRightDifference(nums))
}

func leftRightDifference(nums []int) []int {

	ans := []int{}
	for i := 0; i < len(nums); i++ {
		leftSum, rightSum := 0, 0
		j := 0
		for j < len(nums) {
			if j < i {
				leftSum = leftSum + nums[j]
			}
			if j > i {
				rightSum = rightSum + nums[j]
			}
			j++
		}
		ans = append(ans, int(math.Abs(float64(leftSum-rightSum))))
	}
	return ans

}
