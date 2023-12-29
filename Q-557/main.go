package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	//Output: "s'teL ekat edoCteeL tsetnoc"
	fmt.Println(reverseWords(s))
}
func reverseWords(s string) string {
	str := strings.Fields(s)
	strbyt := []string{}

	for i := 0; i < len(str); i++ {
		byt := []byte(str[i])
		res := []byte{}
		for j := len(byt) - 1; j >= 0; j-- {
			res = append(res, byt[j])
		}
		strbyt = append(strbyt, string(res))

	}
	resstr := strings.Join(strbyt, " ")

	return resstr
}
