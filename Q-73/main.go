package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	//Output: [[1 0 1] [0 0 0] [1 0 1]]
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	var firstRowZero, firstColZero bool

	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}

	fmt.Println(matrix)
}
