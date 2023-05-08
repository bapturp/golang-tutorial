package main

import "fmt"

/*
Closures can change values outside of the function
*/

// Clojure with private variable
func myFunc() func() {
	c := 0
	return func() {
		c++
		fmt.Println("c:", c)
	}
}

func useFunc(f func(int, int) int, x, y int) {
	fmt.Println("Answer:", f(x, y))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {

	intSum := func(x, y int) int { return x + y }
	fmt.Println("5 + 4 =", intSum(5, 4))

	myNum := 5
	increment := func() { myNum++ } // myNum is modified within the function
	increment()

	fmt.Println("myNum is:", myNum) // myNum has been incremented: myNum == 6

	// passing function to a function
	useFunc(sumVals, 5, 8)

	f := myFunc()
	f()
	f()
}
