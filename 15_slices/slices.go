package main

import (
	"fmt"
	"reflect"
)

/*
Slices are like arrays but can grow

  var name []datatype

  slice := make([]string, size, capacity) //
*/

func main() {
	// using make to create a slice of string of size 4
	slice1 := make([]string, 4)
	slice1[0] = "🍌"
	slice1[1] = "🍍"
	slice1[2] = "🥭"

	// Size of the slices is the size as defined in the definition
	fmt.Println(len(slice1))

	// for range
	for _, v := range slice1 {
		fmt.Println(v)
	}

	// slicing an array
	arr := [5]string{"Avocado", "Apple", "Strawberry", "Kiwi", "Lemon"} // Array
	fmt.Println("Original Array", arr)

	slice2 := arr[0:2] // slicing the array from the 0 index up to 2 index (excluded)
	fmt.Println("Slice [0:2]", slice2)

	slice3 := arr[:3] // from start to 3 (excluded)
	fmt.Println("Slice [:3]", slice3)

	slice4 := arr[2:] // from 2 to the last (included)
	fmt.Println("Slice [2:]", slice4)

	// Changing an array will also change the slice
	arr[0] = "Apricot"
	fmt.Println(slice2) // slice[0] == "Apricot"

	// changing a slice will also chnage the array
	slice2[1] = "Grapes"
	fmt.Println(arr)

	// Append to a slice. Affect other slices and orignal array
	slice2 = append(slice2, "Melon")
	fmt.Println("Orignal Array", arr)
	fmt.Println("slice3 ", slice3)

	slice5 := make([]string, 4)
	fmt.Println(slice5)
	fmt.Println(reflect.TypeOf(slice5[0]))

}
