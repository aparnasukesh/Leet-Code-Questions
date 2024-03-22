package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	//Output: 1
	fmt.Println(maxWidthOfVerticalArea(points))
}
func maxWidthOfVerticalArea(points [][]int) int {
	fmt.Println("Points:", points)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	fmt.Println("Points:", points)

	maxWidth := 0

	for i := 1; i < len(points); i++ {
		width := points[i][0] - points[i-1][0]
		fmt.Println("i:", i, width)
		if maxWidth < width {
			maxWidth = width
		}
	}
	return maxWidth
}
