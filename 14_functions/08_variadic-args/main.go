package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // alice of float64
	n := average(data...)                     // add one item at a time
	fmt.Println(n)
}

func average(sf ...float64) float64 { // I have to pass in one or more float64, but not a whole list
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// variadic arg == x...
// variadic params == ...x
