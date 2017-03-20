package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt is a custom square function
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number") // errors pkg and function New
	}
	// implementation
	return 42, nil
}
