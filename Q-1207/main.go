package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))

}
func uniqueOccurrences(arr []int) bool {
	visited := make([]int, len(arr))
	var res []int
	for i := 0; i < len(arr); i++ {

		if visited[i] != -1 {
			count := 0
			for j := 0; j < len(arr); j++ {

				if arr[i] == arr[j] {
					count = count + 1
					visited[i] = -1
					visited[j] = -1
				}
			}
			res = append(res, count)
		}
	}
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res); j++ {
			if res[i] == res[j] && i != j {
				return false
			}
		}
	}
	return true

}
