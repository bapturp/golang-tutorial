package main

import "fmt"

/*
Generics are mainly use for function that accepts multiple datatypes.
*/

/*
Generic function to handle mutiple types

SumIntsOrFloats is a function with:
Two type parameters (inside the brackets): K and V
One argument: m
It returns a value of type V
K type parameter specifies the type constraint `comparable`. comparable are any type whose value can be compared with == or != (Go map requires a key be comparable)
V type parameter specifies a constraint union of two types: int64 and float64.
*/
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

/*
The constraint can be moved into its own interface so that it can be reuse in multiple places.
Declaring constraint in this way helps streamline code, such as when a constraint is more complex
*/
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	m1 := map[int]int64{1: 5, 2: 4}
	m2 := map[int]float64{1: 4.6, 2: 4.7}

	fmt.Println("Calling SumIntsOrFloats()")
	fmt.Println("5 + 4 =", SumIntsOrFloats(m1))
	fmt.Println("5.6 + 4.7 =", SumIntsOrFloats(m2))

	fmt.Println("Calling SumIntsOrFloats()")
	fmt.Println("5 + 4 =", SumNumbers(m1))
	fmt.Println("5 + 4 =", SumNumbers(m2))
}
