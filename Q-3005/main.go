package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 2, 2, 3, 1, 4}
	nums := []int{10, 12, 11, 9, 6, 19, 11}
	//Output: 4
	//Output: 2
	fmt.Println(maxFrequencyElements(nums))
}
func maxFrequencyElements(nums []int) int {

	freq := make(map[int]int)
	maxvalue := 0
	for _, val := range nums {
		freq[val]++
		if freq[val] >= maxvalue {
			maxvalue = freq[val]
		}
	}
	count := 0
	for _, val := range freq {
		if val == maxvalue {
			count = count + val
		}
	}

	return count

}
