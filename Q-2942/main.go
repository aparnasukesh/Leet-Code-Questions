package main

import (
	"fmt"
)

func main() {
	words := []string{"leet", "code"}
	x := 'e'
	fmt.Println(findWordsContaining(words, byte(x)))
}
func findWordsContaining(words []string, x byte) []int {
	ch := rune(x)
	res := []int{}
	for i := 0; i < len(words); i++ {
		flag := 0
		for _, val := range words[i] {
			if ch == val {
				flag = 1
				break
			}
		}
		if flag == 1 {
			res = append(res, i)

		}
	}
	return res
}
