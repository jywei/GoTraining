package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swan: ")
	fmt.Scan(&meters) // take standard input
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards. ")
}
