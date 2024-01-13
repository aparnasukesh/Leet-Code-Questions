package main

import "fmt"

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	//Output: 6
	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) int {
	wealth := 0
	for i := 0; i < len(accounts); i++ {
		wealth1 := 0
		for _, val := range accounts[i] {
			wealth1 = wealth1 + val
		}
		if wealth1 >= wealth {
			wealth = wealth1
		}

	}
	return wealth
}
