package main

import (
	"fmt"
)

/*
Multi-line
comments
*/

// Alias
// allow to write pl("hello") instead of fmt.Println("hello")
var pl = fmt.Println

func main() {
	fmt.Println("hello world")
	pl("hello world")
}
