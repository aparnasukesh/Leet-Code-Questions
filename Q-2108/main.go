package main

import "fmt"

func main() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	//Output:"ada"
	fmt.Println(firstPalindrome(words))
}
func firstPalindrome(words []string) string {
	for _, word := range words {
		if palindrome(word) {
			return word
		}
	}
	return ""
}

func palindrome(word string) bool {
	flag := 0
	for i := 0; i < len(word); i++ {
		if word[i] != word[len(word)-1-i] {
			flag = 1
			break
		}
	}
	if flag == 1 {
		return false
	}
	return true
}
