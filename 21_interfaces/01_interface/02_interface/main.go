package main

import "fmt"

type square struct { // we created a struct (user defined type)
	side float64 // square is a shape, therefore anything having the method is implementing the interface
} // square inplements the shape interface

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64 // so basically interface is a struct for functions and assgined return/input values
	// on the other word, interface is a type that declaire behaviors
	// however, interface type has no method defined attached to it
	// instead, methods are customized by users (with other user defined types)
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	fmt.Printf("%T\n", s)
	fmt.Println(s.area())
	info(s)
}
