package main

import "fmt"

func main() {
	num := 2016
	fmt.Println(checkPerfectNumber(num))
}
func checkPerfectNumber(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i
		}
		if sum == num {
			return true
		}
	}
	return false
}
