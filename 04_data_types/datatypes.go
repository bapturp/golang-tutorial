package main

import (
	"fmt"
	"reflect"
)

func main() {
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""

	fmt.Println(reflect.TypeOf(25))      // int
	fmt.Println(reflect.TypeOf(3.14))    //	float64
	fmt.Println(reflect.TypeOf(true))    // >> bool
	fmt.Println(reflect.TypeOf("Hello")) // >> string
	fmt.Println(reflect.TypeOf('h'))     // >> int32    /!\ it's also a rune

}
