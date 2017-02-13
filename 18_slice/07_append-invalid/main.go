package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	greeting[3] = "Suprabadham" // this one is out of the length, so if you want to add more items in the constrained length, we need to use append()

	// greeting = append(greeting, "Suprabadham")

	fmt.Println(greeting[2])
}
