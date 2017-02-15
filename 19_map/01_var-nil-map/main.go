package main

import "fmt"

func main() {

	var myGreeting map[string]string // NEVER USE THIS!!!
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
	// myGreeting["Tim"] = "Good morning."
	// myGreeting["Jenny"] = "Bonjour."
}

// add these lines:
/*
 */
// and you will get this:
// panic: assignment to entry in nil map

// A nil map is equivalent to an empty map except that no elements may be added
