package main

import (
	"fmt"
	"math"
)

/*
func funcName(parameters) returnType {BODY}
*/

func sum(a int, b int) int {
	return a + b
}

// returning multiple values
func squareProperties(n float64) (float64, float64, float64) {
	// returns perimeter, diameter and area
	return n * 4, math.Sqrt(2) * n, math.Pow(n, 2)
}

// returns an error
func getQuotient(x float64, y float64) (answer float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't devide by zero")
	} else {
		return x / y, nil
	}
}

// veratic function (unknow number of arguments)
func sum2(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(4, 5))

	perimeter, diameter, area := squareProperties(5.0)

	fmt.Printf("Perimeter: %.2f\n", perimeter)
	fmt.Printf("Diameter: %.2f\n", diameter)
	fmt.Printf("Area: %.2f\n", area)

	fmt.Println(getQuotient(5, 0))

	fmt.Println(sum2(1, 2, 3, 4))
}
