package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}
func mostWordsFound(sentences []string) int {
	var length int

	for i := 0; i < len(sentences); i++ {
		words := strings.Split(sentences[i], " ")
		lenWords := len(words)
		if lenWords >= length {
			length = lenWords
		}
	}
	return length

}
