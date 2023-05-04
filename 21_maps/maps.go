package main

import "fmt"

/*
	Map is a collection of key:values

	map[KeyType]ValueType

*/

func main() {
	// the heroes variable is a map of string keys to string values
	// Map are reference types, here heroes doesn't point to an initialized map
	var heroes map[string]string

	// To initialize map, we use the built in make function
	heroes = make(map[string]string)

	// To create and initialize a map it's more common tyo use the following:
	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"

	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}

	fmt.Printf("Batman is %v\n", heroes["Batman"])

	// trying to access a not valid key returns nil characters
	fmt.Println("Chip:", superPets[3])

	_, ok := superPets[3] // returns: value, boolean

	fmt.Println("Is there a 3rd pet:", ok)

	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// delete a key
	delete(heroes, "The Flash")
}
