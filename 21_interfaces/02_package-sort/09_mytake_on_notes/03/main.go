package main

import (
	"fmt"
	"sort"
)

type num []string

func (n num) Len() int           { return len(n) }
func (n num) Swan(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n num) Less(i, j int) bool { return n[i] < n[j] }

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

}
