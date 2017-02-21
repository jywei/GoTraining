package main

import "fmt"

func main() {
	name := "Sydney"
	str, ok := name.(string) // assertion, but fail, coz it's not an interface
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
