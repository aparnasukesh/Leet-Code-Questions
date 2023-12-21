package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
func findMaxConsecutiveOnes(nums []int) int {
	var temp, count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count = 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 1 {
					count = count + 1
				}
				if nums[j] == 0 {
					break
				}
			}
		}
		if temp <= count {
			temp = count
		}
	}
	return temp
}
