package main

import "fmt"

type Teaspoon float64

type Tablespoon float64

type Milliliters float64

// Aassociated Methods
func (t Teaspoon) toMLs() Milliliters {
	return Milliliters(t * 4.92)
}

func (t Tablespoon) toMLs() Milliliters {
	return Milliliters(t * 14.79)
}

func main() {
	teaspoon1 := Teaspoon(18.2)
	fmt.Printf("%.2f teaspon = %.2f Milliliters\n", teaspoon1, teaspoon1.toMLs())

	var tablespoon1 Tablespoon
	tablespoon1 = 5.1
	fmt.Printf("%.2f tablespoon = %.2f Milliliters\n", tablespoon1, tablespoon1.toMLs())

}
