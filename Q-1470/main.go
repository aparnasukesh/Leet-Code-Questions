package main

import "fmt"

func main() {

	nums := []int{2, 5, 1, 3, 4, 7}
	l := len(nums)
	n := l / 2
	result := shuffle(nums, n)
	fmt.Println(result)

}
func shuffle(nums []int, n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}
	return res

}
