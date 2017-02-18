package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string { // any values we created as the type 'person' will be able to use fullName method
	return p.first + p.last
}

// (p person) is the "receiver"
// it is another parameter
// not idiomatic to use "this" or "self"

func main() {
	p1 := person{"James", "Bond", 35}
	p2 := person{"Miss", "Moneypenny", 26}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
