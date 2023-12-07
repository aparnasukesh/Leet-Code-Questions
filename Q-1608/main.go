package main

import "fmt"

func main() {

	arr := []int{0, 4, 0, 5, 0, 3}

	size := len(arr)

	var temp, i, count int

	for i = 1; i <= size; i++ {
		count = 0
		for j := 0; j < size; j++ {
			if i <= arr[j] {
				count = count + 1
			}
		}
		if i == count {
			temp = count
		}
	}
	if temp != 0 {
		fmt.Println(temp)
	} else {
		fmt.Println(-1)
	}

}
