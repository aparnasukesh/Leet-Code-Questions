package main

import "fmt"

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
}

func wiggleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 && nums[i] == 1 {
			i++
			fmt.Println(nums[i])
		} else {
			for j := i + 1; j < len(nums); j++ {
				if j%2 == 0 && nums[j] == 1 {
					j++
				} else {
					temp := nums[i]
					nums[i] = nums[j]
					nums[j] = temp
				}
			}
		}
	}
	fmt.Println(nums)
}
