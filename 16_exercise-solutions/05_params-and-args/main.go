package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...) // pass in multiple values
	foo()
}

func foo(numbers ...int) { // the input parameters will be more than one
	fmt.Println(numbers)
}
