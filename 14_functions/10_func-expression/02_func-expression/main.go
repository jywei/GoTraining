package main

import "fmt"

func main() {
	// if you assign a function to a variable, it's called func expression
	greeting := func() { // anonymous function
		fmt.Println("Hello World!")
	}

	greeting()
}
