package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string) // the first string is the key type, and the second string is the value type
	// var myGreeting = map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}
