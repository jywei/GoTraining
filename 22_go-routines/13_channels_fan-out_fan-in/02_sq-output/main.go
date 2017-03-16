package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// FAN OUT => multiple funcs pulling out from the same channel until it's closed
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN => multiple channels writing to the same channel
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func gen(nums ...int) chan int { // gen takes variadic nums, and output a bidirectional channel
	fmt.Printf("Type of nums %T\n", nums) // here you will get a slice of int

	out := make(chan int)
	go func() {
		for _, n := range nums { // range over a slice of map
			out <- n
		}
		close(out)
	}()
	return out // return the channel
}

func sq(in chan int) chan int { // sq will take one channel and return another channel
	out := make(chan int)
	go func() {
		for n := range in { // this just simply ranges over a channel
			out <- n * n // a channel is a sequential serial pipeline (one at a time)
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	fmt.Printf("Type of cs: %T\n", cs) // here you will get a slice of channels

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) { // anomynous func
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c) // here is to call the function and also a given parameter
	} // if one don't take in c, there will have two go func accessing the same c

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed
FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.
PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.
source:
https://blog.golang.org/pipelines
*/

/*
CHALLENGE #1
When know HOW to do fan out / fan in, but do we know WHY?
Why would we want to do fan out / fan in?
*/

// when we have a lot of values to deal with, we can have multiple workers to process them
