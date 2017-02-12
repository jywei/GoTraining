package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // it runs before main function
	hello()
}
