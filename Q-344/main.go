package main

import "fmt"

func main() {
	input := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString(input)
}

func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
	fmt.Println(s)
}
