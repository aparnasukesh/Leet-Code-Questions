package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 5
	fmt.Println(findComplement(num))
}
func findComplement(num int) int {
	binaryNum := fmt.Sprintf("%b", num)
	compliment := ""
	for i, _ := range binaryNum {
		if binaryNum[i] == '1' {
			compliment = compliment + "0"
		} else {
			compliment = compliment + "1"
		}
	}

	fmt.Println(binaryNum)
	fmt.Println(compliment)
	numInt, _ := strconv.ParseInt(compliment, 2, 64)

	return int(numInt)
}
