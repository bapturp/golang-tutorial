package main

import "fmt"

// https://pkg.go.dev/fmt

func main() {
	// Print Format

	// %s -> string
	// %U -> Unicode format i.e. `U+0073`
	// %c -> character
	s := "Alice"
	fmt.Printf("String/Char/Unicode: %s %c %U\n", s, s[0], s[0])

	// %d -> int/uint base 10
	// %b -> base 2
	// %x / %X -> base 16 lowercase & uppercase letters
	n := 5
	fmt.Printf("Int/Uint: %d %b %x\n", n, n, n)

	// %f -> float
	// %.2f -> default width, precision 2
	f := 3.14
	fmt.Printf("Float: %f %.2f\n", f, f)

	// %t -> boolean
	b := true
	fmt.Printf("Bool: %t\n", b)

	// %v -> value in default format
	fmt.Printf("Default format: %v %v %v %v\n", s, n, f, b)

	// %T -> type of a value
	fmt.Printf("Type: %T %T %T %T\n", s, n, f, b)
}
