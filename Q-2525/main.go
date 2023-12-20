package main

import (
	"fmt"
	"math"
)

func main() {
	length := 1000
	width := 35
	height := 700
	mass := 300
	fmt.Println(categorizeBox(length, width, height, mass))
}
func categorizeBox(length int, width int, height int, mass int) string {
	power := math.Pow(10, 4)
	power1 := math.Pow(10, 9)
	Bulky := 0
	Heavy := 0
	volume := length * width * height
	if (length >= int(power) || width >= int(power) || height >= int(power) || mass >= int(power)) || volume >= int(power1) {
		Bulky = 1
	}
	if mass >= 100 {
		Heavy = 1
	}

	if Bulky == 1 && Heavy == 1 {
		return "Both"
	}
	if Bulky == 0 && Heavy == 0 {
		return "Neither"
	}
	if Bulky == 1 && Heavy == 0 {
		return "Bulky"
	}
	if Bulky == 0 && Heavy == 1 {
		return "Heavy"
	}
	return ""
}
