package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 { // take in the whole list the same time
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
