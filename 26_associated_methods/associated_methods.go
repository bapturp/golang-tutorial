package main

import (
	"fmt"
)

/*
Go does not have classes. However we can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in itsown argument list between the func keyword and the method name.
*/

type Teaspoon float64

type Tablespoon float64

type Milliliters float64

// `toMLs` is a method that has a receiver of type `Teaspoon` name `t`
func (t Teaspoon) toMLs() Milliliters {
	return Milliliters(t * 4.92)
}

// `toMLs` is a method that has a receiver of type `Tablespoon` name `t`
func (t Tablespoon) toMLs() Milliliters {
	return Milliliters(t * 14.79)
}

func main() {
	// teaspoon1 is of type main.Teaspoon
	teaspoon1 := Teaspoon(18.2)

	// calling the method `.toMLs()`
	fmt.Printf("%.2f teaspon = %.2f Milliliters\n", teaspoon1, teaspoon1.toMLs())

	var tablespoon1 Tablespoon
	tablespoon1 = 5.1
	fmt.Printf("%.2f tablespoon = %.2f Milliliters\n", tablespoon1, tablespoon1.toMLs())

}
