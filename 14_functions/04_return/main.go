package main

import "fmt"

func main() {
	fmt.Println(greet("jane ", "Joe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
