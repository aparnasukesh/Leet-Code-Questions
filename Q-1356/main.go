package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arr := []int{10, 100, 1000, 10000}
	//Output: [0,1,2,4,8,3,5,6,7]
	fmt.Println(sortByBits(arr))
}

func sortByBits(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] += count1s(arr[i]) << 14
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] &= (1 << 14) - 1
	}
	return arr
}

func count1s(n int) int {
	c := 0
	for i := 0; i < 14; i++ {
		c += (n >> i) & 1
	}
	return c
}
