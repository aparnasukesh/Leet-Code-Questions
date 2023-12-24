package main

import (
	"fmt"
)

func main() {

	str := "0010"
	fmt.Println(minOperations(str))

}
func minOperations(s string) int {
	byt := []byte(s)
	temp := byt[0]
	count := 0
	fmt.Println(byt)
	for i := 1; i < len(byt); i++ {
		if temp == 48 {
			if i%2 != 0 && byt[i] != 49 {
				byt[i] = 49
				count++
			}
			if i%2 == 0 && byt[i] != 48 {
				byt[i] = 48
				count++
			}
		}
		if temp == 49 {
			if i%2 != 0 && byt[i] != 48 {
				byt[i] = 48
				count++
			}
			if i%2 == 0 && byt[i] != 49 {
				byt[i] = 49
				count++
			}
		}
	}
	return count
}
