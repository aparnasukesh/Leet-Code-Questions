package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	m1 := make(map[int]int)

	for _, val := range nums {
		m1[val]++

	}
	fmt.Println(m1)
	for i, _ := range m1 {
		if m1[i] == 1 {
			return i
		}
	}
	return 0

}
