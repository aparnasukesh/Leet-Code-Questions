package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 0, 1, 1, 9, 7}
	a := 7
	b := 2
	c := 3
	//Output: 4
	fmt.Println(countGoodTriplets(arr, a, b, c))
}
func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				diff1 := math.Abs(float64(arr[i]) - float64(arr[j]))
				diff2 := math.Abs(float64(arr[j]) - float64(arr[k]))
				diff3 := math.Abs(float64(arr[i]) - float64(arr[k]))
				if int(diff1) <= a && int(diff2) <= b && int(diff3) <= c {
					count++
				}
			}
		}

	}
	return count
}
