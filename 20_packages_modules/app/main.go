package main

import (
	"fmt"
	utils "myapp/utils"
)

func main() {
	fmt.Println("Hello", utils.Name)

	arr := []int{1, 2, 3, 4}

	s := utils.IntArrToStrArr(arr)

	for _, str := range s {
		fmt.Println(str)
	}
}
