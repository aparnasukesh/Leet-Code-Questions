package main

import (
	"fmt"
	"math"
)

func main() {
	c := 7
	//Output: true
	//Explanation: 1 * 1 + 2 * 2 = 5
	fmt.Println(judgeSquareSum(c))
}
func judgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))
	fmt.Println(left, right)

	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}
