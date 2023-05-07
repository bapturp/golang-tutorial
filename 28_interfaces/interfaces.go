package main

import "fmt"

/*
Interfaces allow to create contracts, if anything inherits it, it must implement
some defined methods.
*/

// Anything that inherit from Animal must implement the methods AngrySound and HappySound
type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	fmt.Println("Cat attacks its prey")
}

func (c Cat) Name() string {
	return string(c)
}

// As soon as we implement the method AngrySound, Cat inheric from Animal automatically
// We now need to implement all other methods defined in the Animal interface
func (c Cat) AngrySound() {
	fmt.Println("Cat says Hisss")
}

func (c Cat) HappySound() {
	fmt.Println("Cat says Purrrrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Loopy")
	kitty.AngrySound()

	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	fmt.Println(kitty2.Name())
}
