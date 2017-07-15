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

type human interface { // If you have a speak method you are using a human interface
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}
	// p1.speak()
	saySomething(p1) // polymorphism
	// fmt.Println(p1)

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	// sa1.speak()
	// sa1.person.speak()
	saySomething(sa1)
	// fmt.Println(sa1)
}
