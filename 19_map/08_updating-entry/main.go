package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   " Good morning!",
		"Jenny": " Bonjour!",
	}

	myGreeting["Harleen"] = " Howdy"
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = " Gidday"
	fmt.Println(myGreeting)
}

// under a slice is an array
// under a map is a hash table
