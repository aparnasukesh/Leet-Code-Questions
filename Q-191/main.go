package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num uint32 = 00000000000000000000000010000000

	fmt.Println(hammingWeight(num))
}
func hammingWeight(num uint32) int {
	fmt.Println(num)
	count := 0
	strNum := strconv.FormatUint(uint64(num), 2)
	fmt.Println(strNum)
	for _, nums := range strNum {
		if nums == '1' {
			count = count + 1
		}
	}
	return count
}
