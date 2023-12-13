package main

import (
	"fmt"
)

func main() {
	n := 43046720
	res := isPowerOfFour(n)
	fmt.Println(res)
}
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n%4 == 0 {
		fmt.Println(n)
		n = n / 4
	}
	return n == 1
}
