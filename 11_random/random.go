package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seedSeconds := time.Now().Unix()
	rand.Seed(seedSeconds)
	randomNumber := rand.Intn(50) // generate number from 0 to 49
	fmt.Println(randomNumber)
}
