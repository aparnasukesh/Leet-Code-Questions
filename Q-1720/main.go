package main

import (
	"fmt"
)

func main() {
	encoded := []int{1, 2, 3}
	first := 1
	//Output: [1,0,2,1]
	fmt.Println(decode(encoded, first))
}
func decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)
	arr[0] = first
	for i := 0; i < len(encoded); i++ {
		arr[i+1] = encoded[i] ^ arr[i]

	}
	return arr

}
