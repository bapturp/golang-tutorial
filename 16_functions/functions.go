package main

import (
	"fmt"
	"math"
)

// Basic function
func sum(a int, b int) int {
	return a + b
}

// Returns multiple values
func squareProperties(n float64) (float64, float64, float64) {
	// returns perimeter, diameter and area
	return n * 4, math.Sqrt(2) * n, math.Pow(n, 2)
}

// Returns result and error (or nil)
func quotient(x float64, y float64) (answer float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't devide by zero")
	} else {
		return x / y, nil
	}
}

// Variadic function (unknown number of arguments)
// ...T represents any number of aguments of type T
// The actual type inside the function is []T (Array)
func sumVariadic(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// Accept and array as parameter
func arraySum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

// variables are passed as value (as opposed as passed by reference)
// n here is a copy of n as created below.
func changeVal(n int) int {
	n++
	return n
}

func main() {
	// basic function
	result1 := sum(4, 5)
	fmt.Println(result1)

	// returns multiple values
	perimeter, diameter, area := squareProperties(5.0)
	fmt.Printf("Perimeter: %.2f\nDiameter: %.2f\nArea: %.2f\n", perimeter, diameter, area)

	// returns result and err
	result2, err := quotient(5, 0)
	fmt.Println(result2, err)

	// accept any number or arguments
	result3 := sumVariadic(1, 2, 3, 4)
	fmt.Println(result3)

	// variadic function accepts also a slice unpacked with ...
	slice := []int{1, 2, 3, 4, 5}
	result4 := sumVariadic(slice...)
	fmt.Println(result4)

	// accept an array as argument
	arr := []int{1, 2, 3, 4}
	result5 := arraySum(arr)
	fmt.Println(result5)

	n := 5
	fmt.Printf("n before changeVal: %d\n", n)
	changeVal(n) // n is passed by value
	fmt.Printf("n after changeVal: %d\n", n)
}
