package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[3] = "Fizz"
	m[5] = "Buzz"
	m[8] = "Bazz"

	for i := 0; i < 100; i++ {
		var answer string
		for key, element := range m {
			if i%key == 0 {
				answer = answer + element
			}
		}
		if answer != "" {
			fmt.Println(i, answer)
		}
	}
}
