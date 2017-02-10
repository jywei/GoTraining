package main

import "fmt"

func makeGreeter() func() string { // the type of the returned function will be == func() string
	return func() string {
		return "Hello World!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
