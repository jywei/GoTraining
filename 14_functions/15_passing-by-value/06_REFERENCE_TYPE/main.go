package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ] 1
	changeMe(m)
	fmt.Println(m) // [Roy] 3
}

func changeMe(z []string) {
	z[0] = "Roy"
	fmt.Println(z) // [Roy] 2
}
