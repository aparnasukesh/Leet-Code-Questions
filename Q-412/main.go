package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))
}
func fizzBuzz(n int) []string {

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			ans[i] = "FizzBuzz"
		} else if (i+1)%3 == 0 {
			ans[i] = "Fizz"
		} else if (i+1)%5 == 0 {
			ans[i] = "Buzz"
		} else {
			ans[i] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
