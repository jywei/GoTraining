package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if (i % 2) == 0 {
			continue // meaning skip, continue from the top
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
