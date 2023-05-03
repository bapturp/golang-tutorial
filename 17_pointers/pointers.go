package main

import (
	"fmt"
)

// n *int represents a pointer
func changeVal(n *int) {
	*n++
}

func doubleArrayValues(arr *[4]int) {
	for i := range arr {
		arr[i] *= 2
	}
}

func main() {
	n := 5

	// Create a pointer
	p := &n // same as: var p *int = &n

	// passing the pointer
	changeVal(p)

	fmt.Println(n)

	fmt.Println("p address:", p) // Address in memory
	fmt.Println("p value:", *p)  // * deference operator

	arr := [4]int{1, 2, 3, 4}

	// passing the address of the array
	doubleArrayValues(&arr)

	fmt.Println(arr)
}
