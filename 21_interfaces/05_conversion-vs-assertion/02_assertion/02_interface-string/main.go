package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	// var name interface{} = 12 // invalid (not a string)
	str, ok := name.(string) // it's an interface, so we can use assertion
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
