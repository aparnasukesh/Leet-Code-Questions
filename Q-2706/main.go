package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 3}
	money := 3
	fmt.Println(buyChoco(nums, money))
}
func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if prices[0]+prices[1] <= money {
		return money - (prices[0] + prices[1])
	}
	if prices[0]+prices[1] > money {
		return money
	}
	return 0
}
