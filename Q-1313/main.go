package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums))
	//Output: [2,4,4,4]
}
func decompressRLElist(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i += 2 {
		j := i + 1
		for k := 0; k < nums[i]; k++ {
			res = append(res, nums[j])
		}
	}
	return res
}
