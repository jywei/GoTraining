package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if (1 % 3) == 0 {
			fmt.Println(i)
		}
	}
}
