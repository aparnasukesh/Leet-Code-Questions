package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(getConcatenation(nums))
}
func getConcatenation(nums []int) []int {
	ans := []int{}
	for j := 1; j <= 2; j++ {
		ans = append(ans, nums...)

	}
	return ans
}
