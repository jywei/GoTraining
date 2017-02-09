package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name) // take the address
	fmt.Println("Hello,", name)
}
