package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := incrementor(2)
	var count int

	for n := range c { // pull things out of the channel here
		count++
		fmt.Println(n)
	}

	fmt.Println("Final Count:", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", k) // put things into the channel here
			} // Itoa == int to ASCII
			done <- true
		}(i)
	}

	go func() { // block until done channel is true, and close c channel
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}
