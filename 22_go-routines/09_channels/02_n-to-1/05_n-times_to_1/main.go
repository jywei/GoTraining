package main

import "fmt"

func main() {
	n := 10 // 10 things we want to run
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() { // 10 different go routines
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
