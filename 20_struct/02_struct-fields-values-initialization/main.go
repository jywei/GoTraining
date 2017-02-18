package main

import "fmt"

type person struct {
	first string // field: name type
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 35} // variable of the type
	p2 := person{"Miss", "Moneypenny", 26}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
