package main

import (
	"fmt"
	"math/big"
)

func main() {
	num1 := "3876620623801494171"
	num2 := "6529364523802684779"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	num1Int, success1 := new(big.Int).SetString(num1, 10)
	num2Int, success2 := new(big.Int).SetString(num2, 10)

	if !success1 || !success2 {
		return "Number parsing errors"
	}
	sum := new(big.Int).Add(num1Int, num2Int)
	return sum.String()

}
