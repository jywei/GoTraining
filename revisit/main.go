package main

import "fmt"

var x int // package level scope

func main() {
	x = 7
	// x := 7
	// fmt.Println(x)
	fmt.Printf("%T\n", x)
}
