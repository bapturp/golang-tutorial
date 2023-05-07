package main

import (
	"fmt"
	"log"
	utils "myapp/utils"
)

func main() {
	date1 := utils.Date{}

	err := date1.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	err = date1.SetMonth(3)
	if err != nil {
		log.Fatal(err)
	}

	err = date1.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Date is: %d/%d/%d\n", date1.GetDay(), date1.GetMonth(), date1.GetYear())
}
