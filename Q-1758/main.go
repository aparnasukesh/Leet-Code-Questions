package main

import (
	"fmt"
)

func main() {

	str := "10010100"
	fmt.Println(minOperations(str))

}

// func minOperations(s string) int {
// 	byt := []byte(s)
// 	temp := byt[0]
// 	count := 0
// 	fmt.Println(byt)
// 	for i := 1; i < len(byt); i++ {
// 		if temp == 48 {
// 			if i%2 != 0 && byt[i] != 49 {
// 				byt[i] = 49
// 				count++
// 			}
// 			if i%2 == 0 && byt[i] != 48 {
// 				byt[i] = 48
// 				count++
// 			}
// 		}
// 		if temp == 49 {
// 			if i%2 != 0 && byt[i] != 48 {
// 				byt[i] = 48
// 				fmt.Println(i, byt[i])
// 				count++
// 			}
// 			if i%2 == 0 && byt[i] != 49 {
// 				byt[i] = 49
// 				fmt.Println(i, byt[i])

// 				count++
// 			}
// 		}
// 	}
// 	fmt.Println(byt)
// 	return count
// }

func minOperations(s string) int {
	first, second := 0, 0

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] != '0' {
				first++
			} else {
				second++
			}
		} else {
			if s[i] != '1' {
				first++
			} else {
				second++
			}
		}
	}

	return min(len(s)-first, len(s)-second)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
