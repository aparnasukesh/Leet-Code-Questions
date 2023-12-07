package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 88
	fmt.Println(addDigits(num))

}

func addDigits(num int) int {
	for num >= 10 {
		numstr := strconv.Itoa(num)
		sum := 0

		for _, char := range numstr {
			charint, _ := strconv.Atoi(string(char))
			sum += charint
		}

		num = sum
	}
	return num
}
