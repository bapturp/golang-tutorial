package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Float to Int
	cV1 := 1.5
	cV2 := int(cV1)
	fmt.Println(cV2)

	// String to Integer
	cV3 := "50000000"
	cv4, err := strconv.Atoi(cV3) // Atoi -> ASCII to Integer
	fmt.Println(cv4, err, reflect.TypeOf(cv4))

	// Integer to String
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5) // Itoa -> Integer to ASCII
	fmt.Println(cV6, reflect.TypeOf(cV6))

	// String to Float
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		fmt.Println(cV8, reflect.TypeOf(cV8))
	}

	// Float to String
	cV9 := fmt.Sprintf("%f", 3.14)
	fmt.Println(cV9, reflect.TypeOf(cV9))
}
