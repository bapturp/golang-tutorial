package main

import "fmt"

func main() {
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !

	iAge := 8

	if (iAge >= 1) && (iAge <= 18) {
		fmt.Println("Important birthday")
	} else if (iAge == 21) || (iAge == 50) {
		fmt.Println("Important birthday")
	} else if iAge >= 65 {
		fmt.Println("Important birthday")
	} else {
		fmt.Println("Not important birthday")
	}
}
