package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("5+4=", 5+4)
	fmt.Println("5-4=", 5-4)
	fmt.Println("5*4=", 5*4)
	fmt.Println("5/4=", 5/4) // Integer division == 1
	fmt.Println("5%4=", 5%4)

	// increment
	increment := 1
	increment = increment + 1
	increment++

	fmt.Println(increment)

	// float precision
	fmt.Println(0.11111111111111 + 0.11111111111111)

	// Many math function
	fmt.Println("Max(4, 5)", math.Max(4, 5))
	fmt.Println("Min(4, 5)", math.Min(4, 5))
	fmt.Println("Round(5.5)", math.Round(5.5))
	fmt.Println("Floor(5.5)", math.Floor(5.5))
	fmt.Println("Ceil(5.5)", math.Ceil(5.5))
}
