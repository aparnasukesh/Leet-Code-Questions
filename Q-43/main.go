package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "498828660196"
	num2 := "840477629533"
	multiply(num1, num2)
}
func multiply(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	sum := n1 * n2
	fmt.Println(sum)
	sumstr := strconv.FormatInt(int64(sum), 10)
	return sumstr
}
