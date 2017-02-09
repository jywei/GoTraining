package main

import "fmt"

func main() {
	switch "Mhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medicci":
		fmt.Println("Wassup Meducci")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have no friends? Sad.")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
