package main

import "fmt"

func main() {

	nums := []int{8, 1, 2, 2, 3}
	counts := make([]int, len(nums))
	var count int

	for i := 0; i < len(nums); i++ {
		count = 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count = count + 1
			}
		}
		//counts[i] = append(counts[i], count)
		counts[i] = count
		//fmt.Println("Count of Array elements:", counts[i])
	}
	fmt.Println(counts)
	// for i := 0; i < len(counts); i++ {
	// 	fmt.Printf("%d ", counts[i])
	// }
}
