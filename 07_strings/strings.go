package main

import (
	"fmt"
	"strings"
)

func main() {
	// Strings are just an array of bytes []bytes
	s1 := "A word"

	// Replace
	replacer := strings.NewReplacer("A", "Another")
	s2 := replacer.Replace(s1)
	fmt.Println(s2)

	// Length
	length := len(s2) // returns the string length
	fmt.Println(length)

	// Contains
	result := strings.Contains(s2, "Another") // returns true
	fmt.Println("Contains Another:", result)

	// Index Of
	index := strings.Index(s2, "o") // -> returns first index of matching substring
	fmt.Println(index)

	// Replace all match
	replaced := strings.Replace(s2, "o", "0", -1) // -1 for replacing all matching substring
	fmt.Println(replaced)                         // An0ther w0rd

	// remove all leading and trailing whitespaces
	s3 := "\t \nSome words\n "
	s3 = strings.TrimSpace(s3)
	fmt.Println(s3)

	// split on delimiter
	s4 := "a b c d"
	s5 := strings.Split(s4, " ")
	fmt.Println(s5)

	// lowercase
	s6 := strings.ToLower(s1)
	s7 := strings.ToUpper(s1)
	fmt.Println(s6, s7)

	// Prefix, begin with
	s8 := strings.HasPrefix("tacocat", "taco")
	fmt.Println(s8)

	// Suffix
	s9 := strings.HasSuffix("tacocat", "cat")
	fmt.Println(s9)
}
