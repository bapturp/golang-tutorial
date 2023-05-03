package utils

import "strconv"

/*
Convert an int slice to a string slice
*/

// Capitalize variables are available outside of the scope
var Name string = "Baptiste"

// Capitalize functions are available outside of the scope
func IntArrToStrArr(arr []int) []string {
	s := []string{}

	for _, n := range arr {
		s = append(s, strconv.Itoa(n))
	}

	return s
}
