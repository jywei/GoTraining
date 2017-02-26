package main

import "fmt"

func main() {
	c := make(chan int)
	go func() { // the goroutine is runign on its own
		c <- 1
	}()
	fmt.Println(<-c) // so we are blocked here and waiting to receive the <-c
}
