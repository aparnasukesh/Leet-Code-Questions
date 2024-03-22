package main

import "fmt"

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	//Output: "leetcode"

	fmt.Println(restoreString(s, indices))
}

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i := 0; i < len(indices); i++ {
		index := indices[i]
		res[index] = byte(s[i])
	}

	return string(res)
}
