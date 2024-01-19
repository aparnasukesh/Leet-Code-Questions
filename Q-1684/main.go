package main

import (
	"fmt"
)

func main() {
	// allowed := "ab"
	// words := []string{"ad", "bd", "aaab", "baa", "badab"}
	// //Output:2
	allowed := "abc"
	words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println(countConsistentStrings(allowed, words))
}
func countConsistentStrings(allowed string, words []string) int {
	count := 0
	var flag int
	for i := 0; i < len(words); i++ {
		for _, ch := range words[i] {
			flag = 0
			for _, val := range allowed {
				if ch == val {
					flag = 1
				} else {
					flag = 0
				}
			}
		}
		if flag == 1 {
			count++
		}
	}
	return count
}
