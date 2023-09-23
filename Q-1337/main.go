package main

import (
	"fmt"
	"sort"
)

func main() {

	slice1 := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}

	UsingBuitInFunc()

	result := kWeakestRows(slice1, 3)
	fmt.Println("Weakest rows: ", result)

	counts := make(map[int]int)

	for i, row := range slice1 {
		count := 0
		for _, val := range row {
			if val == 1 {
				count++
			}
		}
		counts[i] = count
	}

	var row_index []int
	for len(row_index) < 3 && len(counts) > 0 {
		minIndex, minValue := -1, len(slice1[0])+1
		for i, count := range counts {

			if count < minValue {
				minIndex, minValue = i, count
			}
		}
		row_index = append(row_index, minIndex)
		delete(counts, minIndex)
	}
	fmt.Println("Index of the samllest elements up to 3")
	for _, index := range row_index {
		fmt.Println(index)
	}
}

// The same program using built in function------------------------------------------------------------------

func UsingBuitInFunc() {
	slice1 := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}

	l := len(slice1)
	new_slice := make([]int, l)

	for i := 0; i < l; i++ {
		count := 0
		for _, j := range slice1[i] {
			if j == 1 {
				count++
			}
		}
		new_slice[i] = count
	}
	row_index := make([]int, l)
	for i := range row_index {
		row_index[i] = i
	}
	sort.Slice(row_index, func(i, j int) bool {
		return new_slice[row_index[i]] < new_slice[row_index[j]]
	})

	fmt.Println("Index of the smallest elements up to 3")
	for _, index := range row_index[:3] {
		fmt.Println(index)
	}
}

// Another Method

type mapS struct {
	index int
	val   int
}

func kWeakestRows(mat [][]int, k int) []int {
	result := make([]int, k)
	set := make([]mapS, len(mat))

	for i := 0; i < len(mat); i++ {
		count := 0
		for _, ch := range mat[i] {
			if ch == 1 {
				count++
			}
		}
		set[i].index = i
		set[i].val = count
	}
	sortmap(set)

	for i := 0; i < k; i++ {
		result[i] = set[i].index
	}
	return result
}

func sortmap(set []mapS) {
	for i := 0; i < len(set); i++ {
		for j := i + 1; j < len(set); j++ {
			if set[i].val > set[j].val {
				set[i], set[j] = set[j], set[i]
			}
		}
	}

	sort.Slice(set, func(i, j int) bool {
		if set[i].val != set[j].val {
			return set[i].val < set[j].val
		}
		return set[i].index < set[j].index
	})
}
