package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// create and open file
	file, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	// close the file when we're out of the scope
	defer file.Close()

	prime := []int{1, 3, 5, 7, 11}

	var prime2 []string

	for _, n := range prime {
		prime2 = append(prime2, strconv.Itoa(n))
	}

	for _, num := range prime2 {
		// write to a file
		_, err := file.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the file
	file, err = os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read a file
	scan1 := bufio.NewScanner(file)

	// loop through the Scanner
	for scan1.Scan() {
		fmt.Println("Prime:", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
