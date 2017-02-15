package main

import "fmt"

func main() {
	buckets := make([]int, 1) // bucket is defined as a slice
	fmt.Println(buckets[0])   // 0
	buckets[0] = 42
	fmt.Println(buckets[0]) // 42
	buckets[0]++
	fmt.Println(buckets[0]) // 43
}
