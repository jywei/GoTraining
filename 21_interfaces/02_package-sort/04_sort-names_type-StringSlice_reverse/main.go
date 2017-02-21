package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	// Reverse will take an Interface and return an Interface
	fmt.Println(s)

	// for experimentation to understand what's going on:
	// uncomment and experiment with the code below:

	// sort.Sort(sort.StringSlice(s))
	// fmt.Println(s)
	// fmt.Printf("just s: %T\n", s)

	// s = sort.StringSlice(s)
	// fmt.Printf("just s: %T\n", s) // if you don't reassign s, s will only be a []string
	// t := sort.StringSlice(s)
	// fmt.Printf("just t: %T\n", t) // reassign, a new type

	// fmt.Printf("s converted to StringSlice: %T\n", sort.StringSlice(s))
	// // fmt.Printf("s sorted: %T\n", sort.Sort(sort.StringSlice(s)))
	// fmt.Printf("s reversed: %T\n", sort.Reverse(sort.StringSlice(s)))
	// i := sort.Reverse(sort.StringSlice(s))
	// fmt.Println(i)
	// fmt.Printf("%T\n", i)
	// sort.Sort(i)
	// fmt.Println(s)
}

// https://golang.org/pkg/sort/#Sort
