package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11} // both length and capacoty are 6
	// mySlice := make([]int, 6, 6)

	for i, value := range mySlice {
		fmt.Println(i, " - ", value)
	}
}
