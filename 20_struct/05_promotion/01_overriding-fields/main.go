package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person               // the type is taken from person struct, so the inner type is promoted to the outter type
	First         string // therefore, the creating variable next will get the properpty from both 'person' and 'doublezero'
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven", // the outter type First will ocerwrite inner type First
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.person.First, p1.Last) // Last got promoted
	fmt.Println(p2.First, p2.person.First)
}
