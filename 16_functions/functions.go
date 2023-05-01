package main

import "fmt"

/*
func funcName(parameters) returnType {BODY}
*/
func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(sum(4, 5))
}
