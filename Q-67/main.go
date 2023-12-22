package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))

}
func addBinary(a string, b string) string {
	aint := new(big.Int)
	bint := new(big.Int)
	aint, _ = aint.SetString(a, 2)
	bint, _ = bint.SetString(b, 2)
	resint := new(big.Int).Add(aint, bint)
	return resint.Text(2)
}
