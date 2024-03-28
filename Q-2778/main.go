package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	//Output:21
	fmt.Println(sumOfSquares(nums))
}
func sumOfSquares(nums []int) int {
	length := len(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		if length%(i+1) == 0 {
			sum = sum + (nums[i] * nums[i])
		}
	}
	return sum
}
