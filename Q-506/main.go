package main

import "fmt"

func main() {
	score := []int{5, 4, 3, 2, 1}
	//Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
	fmt.Println(findRelativeRanks(score))

}

func findRelativeRanks(score []int) []string {
	var output []string
	j := 0
	for i := 0; i < len(score); i++ {
		large := score[i]
		for j = 0; j < len(score); j++ {
			if i != j && large <= score[j] && score[j] != -1 {
				large = score[j]

			}
		}
		score[j] = -1
		fmt.Println(large)
	}
	fmt.Println(output)
	return nil
}
