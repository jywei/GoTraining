package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough // go for the next case
	case "Medicci":
		fmt.Println("Wassup Medicci")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Daria":
		fmt.Println("Wassup Daria")
	}
}
