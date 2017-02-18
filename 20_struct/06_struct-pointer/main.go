package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 35} // assign the address of a person struct
	fmt.Println(p1)            // &{James 35}
	fmt.Printf("%T\n", p1)     // *main.person
	fmt.Println(p1.name)       // James
	fmt.Println(*p1)           // address of p1 == {James 35}
	fmt.Println(p1.age)        // 35
}
