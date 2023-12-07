package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{13, 25, 83, 77}

	fmt.Println(separateDigits(nums))
}

func separateDigits(nums []int) []int {
	digits := []int{}
	for _, num := range nums {
		numString := strconv.Itoa(num)
		for _, char := range numString {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}
			digits = append(digits, digit)
		}
	}

	return digits
}
