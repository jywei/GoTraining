package main

import (
	"fmt"
	"time"
)

func main() {
	// channel is like map, so we use make for 'maps', 'slices', and 'channels'
	c := make(chan int) // the unbuffered channel we will communicate with int
	// Unbuffered channel is the channel don't explicitly assign size ie. cha := make(chan int, 10) => cha is a buffered channel
	// two independently running go-routines
	go func() {
		for i := 0; i < 10; i++ {
			c <- i // put i on the channel, then the code blocks
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // we take the number out of the channel, then the code continues
		}
	}()

	time.Sleep(time.Second) // this will give time to go-routines to execute
}

// go run -race main.go
