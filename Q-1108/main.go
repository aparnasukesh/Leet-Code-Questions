package main

import "fmt"

func main() {
	address := "1.1.1.1"
	//Output:"1[.]1[.]1[.]1"
	a := "[.]"
	str := ""
	for _, ch := range address {
		if ch == '.' {
			str = str + a
		} else {
			str = str + string(ch)
		}
	}
	fmt.Println(str)
}
