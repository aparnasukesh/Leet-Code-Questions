package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Printf("enter the number")
	fmt.Scanf("%d", &n)
	if n == 0 {
		return
	}

	size := n + 1

	nums := make([]int, size)

	nums[0] = 0
	nums[1] = 1
	max := -1
	max1 := float64(max)
	for i := 0; i < size; i++ {
		if 2*i >= 2 && 2*i <= n {
			nums[2*i] = nums[i]
		}

		if 2*i+1 >= 2 && 2*i+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
		}
	}

	fmt.Println(nums)
	for _, numbers := range nums {
		max1 = math.Max(max1, float64(numbers))
	}
	max = int(max1)
	fmt.Printf("The Max element in array is : %d", max)
}
