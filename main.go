package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "jsmith@hotmail.com",
			zipCode: 231445,
		},
	}

	fmt.Printf("%+v", john)
}
