package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello how are you Contestant"
	k := 4
	//Output: "Hello how are you"

	fmt.Println(truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	arr := strings.Fields(s)
	str := []string{}
	for i := 0; i < k; i++ {
		str = append(str, arr[i])
	}

	return strings.Join(str, " ")
}
