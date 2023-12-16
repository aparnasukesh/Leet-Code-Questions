package main

import "fmt"

func main() {
	nums :=
		totalHammingDistance()
}

func totalHammingDistance(nums []int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		x := fmt.Sprintf("%064b", nums[i])
		for j := i + 1; j < len(nums); j++ {
			y := fmt.Sprintf("%064b", nums[j])
			for i, _ := range x {
				if x[i] != y[i] {
					count = count + 1
				}
			}
		}
	}
	return count
}
