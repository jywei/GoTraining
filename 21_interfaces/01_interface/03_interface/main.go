package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
} // a concrete type == what it is and what you can do with it

type shape interface { // interface allows things to connect together, like lego blocks
	area() float64
} // an abstract type == no expression or internal value
// square and circle satisfied the shape interface

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
