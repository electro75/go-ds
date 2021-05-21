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

	//struct
	john := person{
		firstName: "John",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "jsmith@hotmail.com",
			zipCode: 231445,
		},
	}

	john.updateName("jim")
	// the above gives the same result as:
	// johnPointer := &john
	// johnPointer.updateName("jim")
	john.print()

	//map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#0000ff",
	}
	colors["white"] = "#fff"
	delete(colors, "white")
	printMap(colors)

	// interface
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// *person can accept a pointer of type person
// it can also accept the type person
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is ", hex)
	}
}
