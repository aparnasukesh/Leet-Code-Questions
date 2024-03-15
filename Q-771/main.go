package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	//Output: 3
	fmt.Println(numJewelsInStones(jewels, stones))
}
func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for i := 0; i < len(jewels); i++ {
		for j := 0; j < len(stones); j++ {
			if jewels[i] == stones[j] {
				count++
			}
		}
	}
	return count
}
