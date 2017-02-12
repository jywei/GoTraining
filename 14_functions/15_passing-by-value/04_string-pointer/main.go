package main

import "fmt"

func main() {

	name := "Roy"
	fmt.Println(&name) // 0x82023c080

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  //Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Roy
	*z = "Rocky"
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Rocky
}
