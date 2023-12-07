package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "0101"
	fmt.Println(minOperations(str))

}
func minOperations(s string) int {

	var num int64
	var err error
	num, err = strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(num); i++ {

	}
}
