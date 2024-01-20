package main

import "fmt"

func main() {
	batteryPercentages := []int{1, 1, 2, 1, 3}
	//Output: 3
	fmt.Println(countTestedDevices(batteryPercentages))
}
func countTestedDevices(batteryPercentages []int) int {
	count := 0
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i] > 0 {
			count++
			for j := i + 1; j < len(batteryPercentages); j++ {
				batteryPercentages[j] = batteryPercentages[j] - 1
			}
		}
	}
	return count
}
