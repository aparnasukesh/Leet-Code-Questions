package main

import "fmt"

func main() {

	num := 121129

	res := isPalindrome(num)
	fmt.Println(res)

	num1 := 121

	res1 := isPalindrome(num1)
	fmt.Println(res1)

}

func isPalindrome(x int) bool {

	num := x

	var res = 0
	var rem int
	if x > 0 {
		for x != 0 {
			rem = x % 10
			res = res*10 + rem
			x = x / 10
		}
	}
	if num == res {
		return true
	}
	return false
}
