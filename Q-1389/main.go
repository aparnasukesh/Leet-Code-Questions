package main

import "fmt"

func main() {

	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}

	target := make([]int, len(nums))
	var j int

	for i := 0; i < len(nums); i++ {
		j = index[i]
		if target[j] == 0 {
			target[j] = nums[i]
		} else {
			for j = len(nums) - 1; j > index[i]; j-- {
				target[j] = target[j-1]
			}

			target[j] = nums[i]
		}
	}
	fmt.Println(target)
}
