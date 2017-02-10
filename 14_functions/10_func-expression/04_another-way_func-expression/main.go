package main

import "fmt"

func makeGreeter() func() string { // makeGreeter will return a function
	return func() string {
		return "Hello World!"
	}
}

func main() {
	greet := makeGreeter() // func expression, so greet() is the function returned from makeGreeter()
	fmt.Println(greet())
}
