package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() { // the outter type will overwrite inner type
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  35,
		},
		LicenseToKill: true,
	}
	p1.Greeting()        // "I'm just a regular person."
	p2.Greeting()        // "Miss Moneypenny, so good to see you."
	p2.person.Greeting() // "I'm just a regular person."
}
