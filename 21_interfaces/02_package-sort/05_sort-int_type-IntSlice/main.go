package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(n)
	sort.Sort(sort.IntSlice(n)) // sort.IntSlice(n) change n into an interface
	fmt.Println(n)
}
