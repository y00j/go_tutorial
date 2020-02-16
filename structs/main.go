package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string7u
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Daly",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 94105,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
