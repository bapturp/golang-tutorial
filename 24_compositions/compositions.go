package main

import "fmt"

type contact struct {
	firstname string
	lastname  string
	phone     string
}

type business struct {
	name    string
	address string
	contact
}

// info is tied to the business struct
func (b business) info() {
	// accessing sub-struct
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.firstname, b.contact.lastname)

}

func main() {
	contact1 := contact{firstname: "Paul", lastname: "Smith", phone: "123-4567"}
	business1 := business{name: "B&S", address: "948 Homer st", contact: contact1}

	business1.info()
}
