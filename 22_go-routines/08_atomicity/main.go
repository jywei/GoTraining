package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64 // it means we are getting ready for atomic stuff

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)                              // you only lock a small thing, very close to mutex
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) // access without race
		// fmt.Println(s, i, "Counter:", counter)                    // still having an unordered problem
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
