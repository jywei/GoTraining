package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	myGreeting1 := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)
	fmt.Println(myGreeting1)
	delete(myGreeting, "two")
	delete(myGreeting1, 2)
	fmt.Println(myGreeting)
	fmt.Println(myGreeting1)
}
