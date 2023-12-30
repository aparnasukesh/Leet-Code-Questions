package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example string in lowercase
	uppercaseString := "Hello, World!"
	lower := "hai all"
	// Convert the string to uppercase
	lowercaseString := strings.ToLower(uppercaseString)
	upper := strings.ToUpper(lower)

	// Print the result
	fmt.Println("Original String:", lowercaseString)
	fmt.Println("Uppercase String:", uppercaseString)

	fmt.Println(lower)
	fmt.Println(upper)
}
