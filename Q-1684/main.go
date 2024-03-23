package main

import (
	"fmt"
)

func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	//Output: 2
	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	allowedMap := make(map[byte]bool)
	for i := range allowed {
		allowedMap[allowed[i]] = true
	}
	for _, word := range words {
		consistent := true
		for j := range word {
			if !allowedMap[word[j]] {
				consistent = false
				break
			}
		}
		if consistent {
			count++
		}
	}
	return count
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	allowed := "ab"
// 	words := []string{"ad", "bd", "aaab", "baa", "badab"}
// 	//Output: 2
// 	fmt.Println(countConsistentStrings(allowed, words))
// }

// func countConsistentStrings(allowed string, words []string) int {
// 	count := 0
// 	for _, word := range words {
// 		consistent := true
// 		for _, char := range word {
// 			if !contains(allowed, byte(char)) {
// 				consistent = false
// 				break
// 			}
// 		}
// 		if consistent {
// 			count++
// 		}
// 	}
// 	return count
// }

// func contains(str string, char byte) bool {
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == char {
// 			return true
// 		}
// 	}
// 	return false
// }
