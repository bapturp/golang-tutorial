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
			break
		}
		c++
	}
}
