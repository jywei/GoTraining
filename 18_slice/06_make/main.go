package main

import "fmt"

func main() {

	customerNumber := make([]int, 3)
	// a := new([100]int)[0:50] == make([]int, 50, 100)

	// 3 is both length & capacity
	// length - number of elements referred to by the slice
	// capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"

	fmt.Println(greeting[2])
}
