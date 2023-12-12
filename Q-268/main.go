package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}

	res := missingNumber(nums)
	fmt.Println(res)
}
func missingNumber(nums []int) int {
	length := len(nums)
	for i := 0; i <= length; i++ {
		flag := 0
		for j := 0; j < length; j++ {
			if i == nums[j] {
				flag = 1
			}
		}
		if flag == 0 {
			return i
		}
	}
	return 0
}
