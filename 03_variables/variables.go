package main

import "fmt"

func main() {
	var vName string = "Baptiste"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"

	// variables are mutables by default as long as the datatype match

	v4 := 2.4
	v4 = 5.4

	fmt.Println(vName, v1, v2, v3, v4)
}
