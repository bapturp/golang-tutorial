package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// In go Characters are called Runes
	// They represent a unicode character

	s1 := "abcdefg"
	fmt.Println("Rune Count:", utf8.RuneCountInString(s1))

	for i, runeVal := range s1 {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal) // %d for Integer %#U for Unicode %c for a character
	}
}
