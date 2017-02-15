package main

import "fmt"

func main() {
	n := hashBucket("Go", 12) // 'G' and the numbers of bucket I want to create
	fmt.Println(n)
}

func hashBucket(word string, buckets int32) int32 { // create a word and a bucket I want to create
	letter := rune(word[0])
	bucket := letter % buckets
	return bucket
}
