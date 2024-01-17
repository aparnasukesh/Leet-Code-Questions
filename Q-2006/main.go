package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	//Output:4
	fmt.Println(countKDifference(nums, k))

}

func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] == k || nums[j]-nums[i] == k {
				count++
			}
		}
	}
	return count
}
