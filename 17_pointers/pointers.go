package main

import "fmt"

// n *int represents a pointer
func changeVal(n *int) {
	*n++
}

func doubleArrayValues(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func average(s ...float64) float64 {
	var r float64
	for _, v := range s {
		r += v
	}
	r /= float64(len(s))
	return r
}

func main() {
	n := 5

	// & address of operator
	// we're passing the address of n
	changeVal(&n)

	fmt.Println(n)

	// creating a pointer that points to n
	var nPtr *int = &n

	fmt.Println("nPtr address:", nPtr) // Address in memory
	fmt.Println("nPtr value:", *nPtr)  // * deference operator

	arr := [4]int{1, 2, 3, 4}
	doubleArrayValues(&arr)
	fmt.Println(arr)

	slice := []float64{5, 11, 19, 23}
	fmt.Printf("Average: %.2f\n", average(slice...))
}
