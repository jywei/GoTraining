package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // you can no longer put values on the channel, but you still still receive values
	}()

	for n := range c { // these codes are blocked until there is a value in c
		fmt.Println(n)
	}
}
