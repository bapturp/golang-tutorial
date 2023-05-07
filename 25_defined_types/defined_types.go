package main

import "fmt"

type Teaspoon float64

type Tablespoon float64

type Milliliters float64

// Don't do that!
func TeaspoonToMl(t Teaspoon) Milliliters {
	return Milliliters(t * 4.92)
}

func TablespoonToMl(t Tablespoon) Milliliters {
	return Milliliters(t * 14.79)
}

// Do associated methods instead (next chapter)

func main() {
	ml1 := Milliliters(Teaspoon(3) * 4.9)
	fmt.Printf("3 teaspoons = %.2f Milliliters\n", ml1)

	ml2 := Milliliters(Tablespoon(2) * 14.79)
	fmt.Printf("2 tablespoons = %.2f Milliliters\n", ml2)

	// Arithmetic
	fmt.Println("2 teaspoons + 4 teaspoons =", Teaspoon(2)+Teaspoon(4))

	// Comparaisons
	fmt.Println("2 teaspoons > 1 teaspoon =", Teaspoon(2) > Teaspoon(1))
}
