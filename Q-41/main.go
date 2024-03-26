package main

import "fmt"

func main() {
	nums := []int{1, 2, 0}
	//Output: 3
	fmt.Println(firstMissingPositive(nums))
}
func firstMissingPositive(nums []int) int {

	res := make(map[int]bool)
	for _, val := range nums {
		res[val] = true
	}
	for i := 1; ; i++ {
		if !res[i] {
			return i
		}
	}
}
