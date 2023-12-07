package main

import "fmt"

func main() {

	arr := []int{9, 9, 3, 9}

	res := IncrementDigits(arr)
	fmt.Println("\nResult Array:", res)
}

func IncrementDigits(digits []int) []int {

	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {

		CurrentSum := digits[i] + carry
		digits[i] = CurrentSum % 10
		carry = CurrentSum / 10

		fmt.Printf("\nCurrentSum: %v\ndigits:%v\ncarry:%v", CurrentSum, digits[i], carry)
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
