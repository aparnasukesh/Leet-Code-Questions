package main

import "fmt"

func main() {
	numRows := 5
	generate(numRows)
}
func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 1; i <= numRows; i++ {
		temp = make([]int, i)
		for j := 1; j <= i; j++ {
			temp[j] = j
		}
	}
	fmt.Println(temp)
	return res
}
