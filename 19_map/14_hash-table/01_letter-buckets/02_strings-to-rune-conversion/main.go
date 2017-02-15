package main

import "fmt"

func main() {
	letter := rune("A"[0]) // strings are a slice of byte, but rune is a byte
	fmt.Println(letter)
}
