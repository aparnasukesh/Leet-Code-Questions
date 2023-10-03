package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 2, 4, 1, 3, 3, 4}
	result := findLucky(arr)
	fmt.Println("The lucky integer is:", result)

	result1 := findLucky1(arr)
	fmt.Println("The lucky integer is:", result1)
}

func findLucky(arr []int) int {
	visited := make([]int, len(arr))
	var count, i int
	var large int
	for i = 0; i < len(arr); i++ {

		if visited[i] != -1 {
			temp := arr[i]
			count = 1

			for j := i + 1; j < len(arr); j++ {
				if temp == arr[j] {
					visited[i] = -1
					visited[j] = -1
					count = count + 1
				}
			}
			if count == temp && count > large {
				large = count

			}
		}
	}
	if large > 0 {
		return large
	}
	return -1
}

// Anotrher method to solve this question

func findLucky1(arr []int) int {
	var large int = -1

	freq := make(map[int]int)
	for _, ch := range arr {
		freq[ch]++
		fmt.Println(freq) // to show iteration
	}
	fmt.Println("The map is :", freq)

	for num, freq := range freq {
		if num == freq && num > large {
			large = num
		}
	}

	return large
}
