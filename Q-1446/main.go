package main

import "fmt"

func main() {
	s := "leetcode"
	fmt.Println(maxPower(s))
}
func maxPower(s string) int {
	var count, temp int
	for i := 0; i < len(s); i++ {
		count = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				count = count + 1
			} else {
				break
			}
		}
		if temp <= count {
			temp = count
		}
	}
	return temp
}
