package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{3, 4, 5, 1, 2}
	//nums := []int{5, 5, 6, 6, 6, 9, 1, 2}
	nums := []int{6, 10, 6}
	fmt.Println(check(nums))
}
func check(nums []int) bool {
	temp := nums[0]
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	var index int
	var rotated []int

	for i := 0; i < len(sorted); i++ {
		if sorted[i] == temp {
			index = i
			break

		}
		//rotated = append(sorted[index:], sorted[:index]...)

	}
	rotated = append(sorted[index:], sorted[:index]...)
	fmt.Println(sorted)
	fmt.Println(rotated)
	if fmt.Sprintf("%v", rotated) == fmt.Sprintf("%v", nums) {
		return true
	}
	return false

}
