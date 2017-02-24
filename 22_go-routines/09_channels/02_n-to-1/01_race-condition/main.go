package main

import (
	"fmt"
	"sync"
)

// To demostrate how to have many functions writing to the same channel

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1) // two go routines tried to access the same variable 'wg'
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
