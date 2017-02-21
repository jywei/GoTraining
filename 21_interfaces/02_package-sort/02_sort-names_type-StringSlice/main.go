package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	//	sort.StringSlice(s).Sort()
	// sort.StringSlice(s) will turn s into a StringSlice type, then we have sort() method
	// len, less, and swap also available for stringslice
	sort.Sort(sort.StringSlice(s)) // sort.Sort() wants to take an Interface interface
	// "sort" is the package, and "Sort" is the function
	fmt.Println(s)
}

// https://golang.org/pkg/sort/#Sort
