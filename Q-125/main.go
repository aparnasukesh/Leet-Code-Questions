package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
func isPalindrome(s string) bool {
	var result string

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result += string(unicode.ToLower(char))
		}
	}
	for i := 0; i < len(result); i++ {
		if result[i] != result[len(result)-i-1] {
			return false
		}
	}
	return true
}

// without using inbuilt function-----------------------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := "A man, a plan, a canal: Panama"

// 	result := removeNonAlphanumericAndLowercase(s)

// 	fmt.Println(result)
// }

// func removeNonAlphanumericAndLowercase(s string) string {
// 	var result string

// 	for _, char := range s {
// 		// Check if the character is alphanumeric.
// 		if isAlphanumeric(char) {
// 			// Convert the character to lowercase.
// 			result += toLowercase(char)
// 		}
// 	}

// 	return result
// }

// func isAlphanumeric(char rune) bool {
// 	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
// }

// func toLowercase(char rune) string {
// 	if char >= 'A' && char <= 'Z' {
// 		// Convert uppercase to lowercase by adding 32 to the ASCII value.
// 		return string(char + 32)
// 	}
// 	return string(char)
// }
