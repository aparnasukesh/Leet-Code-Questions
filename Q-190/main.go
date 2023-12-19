package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := "00000010100101000001111010011100"
	num, _ := strconv.ParseUint(n, 2, 32)
	fmt.Println(uint32(num))
	fmt.Println(reverseBits(uint32(num)))
}
func reverseBits(num uint32) uint32 {
	str := fmt.Sprintf("%032b", num)
	strByte := []byte(str)
	fmt.Println(str)
	for i := 0; i < len(strByte)/2; i++ {

		strByte[i], strByte[len(str)-i-1] = strByte[len(str)-i-1], strByte[i]
	}
	strbytestr := string(strByte)
	res, _ := strconv.ParseUint(strbytestr, 2, 32)

	return uint32(res)
}
