package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
	//Output: [true,true,true,false,true]
}
func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := []bool{}

	for i := 0; i < len(candies); i++ {
		total := candies[i] + extraCandies
		flag := 0
		for j := 0; j < len(candies); j++ {
			if total >= candies[j] {
				flag = 1
			} else {
				flag = 0
				break
			}
		}
		if flag == 0 {
			res = append(res, false)
		}
		if flag == 1 {
			res = append(res, true)
		}
	}
	return res

}
