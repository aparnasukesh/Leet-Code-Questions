package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	//Output:25
	fmt.Println(diagonalSum(mat))
}

func diagonalSum(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if i == j || i+j == len(mat)-1 {
				sum = sum + mat[i][j]
			}
		}
	}
	return sum
}
