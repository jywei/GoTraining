package main

import "fmt"

func main() {
	var x [58]string
	// if there is a number in the bracket, it's an array. If not, it's a slice
	// array is not dynamic, but a slice is

	for i := 65; i <= 122; i++ { // A to z in ASCII
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}
