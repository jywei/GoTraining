package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct { // composition
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}
	p1.speak()
	// fmt.Println(p1)

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()
	// fmt.Println(sa1)
}
