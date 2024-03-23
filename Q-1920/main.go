package main

import "fmt"

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	//Output: [0,1,2,4,5,3]
	fmt.Println(buildArray(nums))
}
func buildArray(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}
	return ans
}
