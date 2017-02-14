package main

import "fmt"

func main() {

	var myGreeting map[string]string // the first string is the key type, and the second string is the value type
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map
