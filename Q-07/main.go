package main

import "fmt"

func main() {
	x := 123
	//Output: 321

	fmt.Println(reverse(x))

}
func reverse(x int) int {
	const MaxInt32 = 1<<31 - 1
	const MinInt32 = -1 << 31
	sum := 0
	fmt.Println(MaxInt32, MinInt32)
	fmt.Printf("%b\n %b:", MaxInt32, MinInt32)
	fmt.Println()
	for x != 0 {
		rem := x % 10
		if sum > MaxInt32/10 || (sum == MaxInt32/10 && rem > 7) {
			return 0
		}
		if sum < MinInt32/10 || (sum == MinInt32/10 && rem < -8) {
			return 0
		}
		sum = (sum * 10) + rem
		x = x / 10
	}
	return sum
}
