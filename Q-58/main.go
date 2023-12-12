package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {

	words := strings.Split(s, " ") // words:=strings.Fields(s) ,we can use this

	fmt.Println(words)
	length := len(words)
	fmt.Println(length)
	num := len(words[length-1])
	return num

}
