package main

import (
	"fmt"
)

func main() {
	s := "011101"
	fmt.Println(maxScore(s))
}
func maxScore(s string) int {
	maxSum := 0
	for i := 1; i < len(s); i++ {
		sum := 0
		count1 := 0
		count2 := 0
		for j := 0; j < len(s); j++ {
			if j < i {
				if s[j] == '0' {
					count1++
				}
			} else {
				if s[j] == '1' {
					count2++
				}
			}
		}

		sum = count1 + count2
		if maxSum <= sum {
			maxSum = sum
		}

	}
	return maxSum
}
