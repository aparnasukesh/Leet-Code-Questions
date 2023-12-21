package main

import "fmt"

func main() {
	s := "111000"
	fmt.Println(checkZeroOnes(s))
}
func checkZeroOnes(s string) bool {
	var temp1, count1, temp2, count2 int
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count1 = 1
			for j := i + 1; j < len(s); j++ {
				if s[j] == '1' {
					count1 = count1 + 1
				}
				if s[j] == '0' {
					break
				}
			}
			if temp1 <= count1 {
				temp1 = count1
			}
		} else {
			count2 = 1
			for j := i + 1; j < len(s); j++ {
				if s[j] == '0' {
					count2 = count2 + 1
				}
				if s[j] == '1' {
					break
				}
			}
			if temp2 <= count2 {
				temp2 = count2
			}

		}
	}
	return temp1 > temp2
}
