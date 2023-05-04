package main

import "fmt"

type customer struct {
	name    string
	address string
	balance float64
}

func getCustomerInfo(c customer) {
	fmt.Printf("%s has $%.2f\n", c.name, c.balance)
}

func newCustomerAddress(c *customer, address string) {
	c.address = address
}

func main() {
	var customer1 customer
	customer1.name = "Tom Smith"
	customer1.address = "1134 Burnaby st"
	customer1.balance = 234.56

	newCustomerAddress(&customer1, "1330 Pendrell st")
	getCustomerInfo(customer1)

	customer2 := customer{"Sally Smith", "7690 Main st", 0.0}
	getCustomerInfo(customer2)
}
