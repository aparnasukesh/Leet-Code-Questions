package main

import "fmt"

func main() {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	//Output: [[1,0,0],[0,1,0],[1,1,1]]
	fmt.Println(flipAndInvertImage(image))
}

func flipAndInvertImage(image [][]int) [][]int {
	res := [][]int{}
	for i := 0; i < len(image); i++ {
		flip := []int{}
		temp := image[i]
		for j := (len(temp) - 1); j >= 0; j-- {
			if temp[j] == 1 {
				flip = append(flip, 0)
			} else {
				flip = append(flip, 1)
			}
		}
		res = append(res, flip)

	}
	return res
}
