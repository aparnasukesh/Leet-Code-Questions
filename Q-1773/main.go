package main

import (
	"fmt"
)

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	fmt.Println(countMatches(items, ruleKey, ruleValue))
	//output=2
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	for i := 0; i < len(items); i++ {
		switch ruleKey {
		case "type":
			if ruleValue == items[i][0] {
				count++
			}
		case "color":
			if ruleValue == items[i][1] {
				count++
			}
		case "name":
			if ruleValue == items[i][2] {
				count++
			}
		}

	}
	return count
}
