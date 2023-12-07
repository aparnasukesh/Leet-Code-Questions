package main

import "fmt"

func main() {

	//nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	nums := []int{1, 1}

	res := findDisappearedNumbers(nums)
	fmt.Println(res)
}

func findDisappearedNumbers(nums []int) []int {

	limit := len(nums)

	var res []int

	for i := 1; i <= limit; i++ {
		count := 0
		for j := 0; j < limit; j++ {

			if i == nums[j] {
				count = count + 1
				break
			}
		}
		if count == 0 {
			res = append(res, i)
		}
	}
	return res
}
