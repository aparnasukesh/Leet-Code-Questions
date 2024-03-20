package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	//Output: 2.00000
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}
	if i >= len(nums1)-1 {
		arr = append(arr, nums2[j:]...)
	}
	if j >= len(nums2)-1 {
		arr = append(arr, nums1[i:]...)
	}
	length := len(arr)
	if length%2 != 0 {
		return float64(arr[(length-1)/2])
	}
	l := float64(arr[(length-1)/2]+arr[length/2]) / 2
	return l
}
