package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c { // if we are going to range it, we have to close it before
		fmt.Println(n)
	}
}

// remember to close your channel
// if you do not close your channel, you will receive this error
// fatal error: all goroutines are asleep - deadlock!
