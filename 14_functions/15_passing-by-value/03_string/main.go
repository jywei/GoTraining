package main

import "fmt"

func main() {

	name := "Roy"
	fmt.Println(name) // Roy

	changeMe(name)

	fmt.Println(name) // Roy
}

func changeMe(z string) {
	fmt.Println(z) // Roy
	z = "Rocky"
	fmt.Println(z) // Rocky
}
