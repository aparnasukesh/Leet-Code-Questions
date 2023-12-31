package main

import "fmt"

func main() {
	numRows := 6
	generate(numRows)
}

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				temp[j] = 1
			} else {
				temp[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res, temp)
	}
	fmt.Println(res)
	return res
}
