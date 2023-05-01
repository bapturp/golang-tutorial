package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seedSeconds := time.Now().Unix()
	rand.Seed(seedSeconds)
	randomNumber := rand.Intn(100) + 1

	for true {
		fmt.Println("Guess a number between 1 and 100")

		// get a buffer reader
		reader := bufio.NewReader(os.Stdin)
		// read the string
		guess, err := reader.ReadString('\n')

		// check for errors while reading from stdin
		if err != nil {
			log.Fatal(err)
		}

		// trim whitespaces
		guess = strings.TrimSpace(guess)

		// cast to Integer
		iGuess, err := strconv.Atoi(guess)
		// check for errors while casting
		if err != nil {
			log.Fatal(err)
		}

		// logic to check if the guessed number is correct
		if iGuess > randomNumber {
			fmt.Println("you're too high!")
		} else if iGuess < randomNumber {
			fmt.Println("you're too low!")
		} else {
			fmt.Println("You found it!")
			break
		}
	}
}
