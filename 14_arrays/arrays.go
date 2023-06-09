package main

import "fmt"

/*
	Arrays have a fixed size and connot be resized

	They come with default values:
	0     -> int
	0.0   -> float
	false -> bool
	""    -> string
*/

func main() {
	// array of size 5 of type int
	var arr1 [5]int

	// with initial values
	// ... specifies a length equal to the number of elements in the literal
	primes := [...]int{2, 3, 5, 7, 11, 13}

	// set values
	arr1[0] = 1

	// array length
	length := len(primes)
	fmt.Printf("primes length %d\n", length)

	// classic for loop
	for i := 0; i < len(primes); i++ {
		fmt.Printf("index %d\t\tvalue: %d\n", i, primes[i])
	}

	// for loop range
	for index, value := range primes {
		fmt.Printf("%d : %d\n", index, value)
	}

	// multi dimentional arrays
	board := [2][2]string{{"🍇", "🍈"}, {"🍉", "🍊"}}
	for _, b := range board {
		for _, v := range b {
			fmt.Println(v)
		}
	}
}
