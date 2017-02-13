package main

import "fmt"

func main() {

	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slice, "c", "g"
	fmt.Println(mySlice[2])    // index access; accessing by index, "c"
	fmt.Println("myString"[2]) // index access; accessing by index, "S" == "83"
}
