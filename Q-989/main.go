package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	k := 516
	res := addToArrayForm(num, k)
	fmt.Println(res)
}

func addToArrayForm(num []int, k int) []int {
	str := ""
	for i := 0; i < len(num); i++ {
		char := strconv.Itoa(num[i])
		str = str + char
	}
	fmt.Println(str)
	strInt, success := new(big.Int).SetString(str, 10)
	if !success {
		return nil
	}
	bigK := big.NewInt(int64(k))
	resNum := new(big.Int).Add(strInt, bigK)
	str = resNum.String()
	res := []int{}
	for _, ch := range str {
		n, _ := strconv.Atoi(string(ch))
		res = append(res, n)
	}
	return res
}
