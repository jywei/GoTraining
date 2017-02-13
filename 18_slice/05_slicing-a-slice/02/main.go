package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Buenos dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:2] ") // 1st
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ") // 1st and 2nd
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ") // 6th and 7th
	fmt.Println(greeting[5:])
	fmt.Print("[:] ") // all
	fmt.Println(greeting[:])
}
