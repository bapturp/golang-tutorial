package main

import "fmt"

func main() {
	// for initialization; condition; post statement {BODY}

	// for loop
	for x := 1; x <= 5; x++ {
		fmt.Println(x)
	}

	// while
	fx := 0
	for fx < 5 {
		fmt.Println(fx)
		fx++
	}

	// infinite
	c := 0
	for true {
		if c > 100 {
			fmt.Println("c:", c)
			break
		}
		c++
	}

	// loop over array using range
	a := []int{1, 2, 3} // array of int with 3 values
	// for index, element of range
	for i, n := range a {
		fmt.Println(i, n)
	}
}
