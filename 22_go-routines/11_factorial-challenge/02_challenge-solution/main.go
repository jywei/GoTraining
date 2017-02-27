package main

import "fmt"

func main() {
	c := factorial(4)
	for n := range c { // range will block from exiting main
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total // when we put a value on the channel, the range will be there waiting to receive
		close(out)
	}()
	// fmt.Println(out) // this will be the address of out channel
	return out
}
