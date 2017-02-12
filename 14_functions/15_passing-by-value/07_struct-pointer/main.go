package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Roy", 22}
	fmt.Println(&c1.name) // 0x8201e4120

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 22}
	fmt.Println(&c1.name) // 0x8201e4120
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Roy 22}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 22}
	fmt.Println(&z.name) // 0x8201e4120

}
