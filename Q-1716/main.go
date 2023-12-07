package main

import "fmt"

func main() {

	n := 30
	TotalAmount := LeetCodeBank(n)
	fmt.Println(TotalAmount)
}

func LeetCodeBank(n int) int {
	var add int = 1
	var res int
	var inc int = 2

	for i := 1; i <= n; i++ {
		res = res + add
		add++

		if i%7 == 0 {
			add = inc
			inc++
		}
	}

	return res
}
