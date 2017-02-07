package main

import "fmt"

var x = 42 // whole package

// var x int = 42

func main() {
	fmt.Println(x)
	y := 17
	foo(y)
	fmt.Println(y) // sequence matters
}

func foo(y int) {
	fmt.Println(x)
	fmt.Println(y)
}

// fixed!
