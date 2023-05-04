package main

import "fmt"

type rectangle struct {
	length, width float64
}

// this function is part of the struct rectangle
func (r rectangle) Perimeter() float64 {
	return r.length*2 + r.width*2
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func main() {
	var r1 rectangle
	r1.length = 10
	r1.width = 5

	// struct litteral
	r2 := rectangle{4.4, 7.9}

	fmt.Printf("r1 perimeter: %.2fm\n", r1.Perimeter())

	fmt.Printf("r2 area: %.2fm2\n", r2.Area())
}
