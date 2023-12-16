package main

import (
	"fmt"
)

func main() {
	res := hammingDistance(1, 4)
	fmt.Println(res)
}

func hammingDistance(x int, y int) int {
	xbinary := fmt.Sprintf("%064b", x)
	ybinary := fmt.Sprintf("%064b", y)
	fmt.Println(xbinary, ybinary)
	count := 0
	for i, _ := range xbinary {
		if xbinary[i] != ybinary[i] {
			count += 1
		}
	}
	return count
}

// func totalHammingDistance(nums []int) int {
//     n := len(nums)
//     totalDistance := 0

//     for i := 0; i < 64; i++ {
//         // Count the number of set bits and unset bits at position i
//         setBits := 0
//         for _, num := range nums {
//             setBits += (num >> i) & 1
//         }
//         unsetBits := n - setBits

//         // Add the contribution to the total Hamming distance
//         totalDistance += setBits * unsetBits
//     }

//     return totalDistance
// }
