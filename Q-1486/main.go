package main

import "fmt"

func main() {
	n := 5
	start := 0
	res := xorOperation(n, start)
	fmt.Println(res)

}
func xorOperation(n int, start int) int {

	var num int
	for i := 0; i < n; i++ {
		num = num ^ (start + 2*i)
	}
	return num

}
