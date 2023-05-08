package main

import (
	"fmt"
	"time"
)

/*
Concurrent tasks in go are called goroutines and not thread.

*/

func printTo15() {
	for i := 0; i <= 15; i++ {
		fmt.Printf("Function 1: %d\n", i)
	}
}

func printTo10() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Function 2: %d\n", i)
	}
}

func main() {
	// go before the function make them a goroutine
	go printTo10()
	go printTo15()

	// the main function doesn't wait for the goroutine to finish
	// One way to deal with it is to pause the main function:
	time.Sleep(2 * time.Second)
}
