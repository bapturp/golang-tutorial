package main

import (
	"fmt"
	"regexp"
)

func main() {
	// str1 := "The ape was at the apex"
	// match, _ := regexp.MatchString("(ape[^ ]?", str1)
	// fmt.Println(match)

	r, _ := regexp.Compile("([crmfp]at)")
	str2 := "Cat rat mat fat pat"

	fmt.Println("MatchString:", r.MatchString(str2))

	fmt.Println("FindString:", r.FindString(str2))
	fmt.Println("Index:", r.FindStringIndex(str2))
	fmt.Println("All String:", r.FindAllString(str2, -1))
	fmt.Println("First 2 Strings:", r.FindAllString(str2, 2))
	fmt.Println("All Submatch Index:", r.FindAllStringSubmatchIndex(str2, -1))
	fmt.Println("Replace all string", r.ReplaceAllString(str2, "dog"))
}
