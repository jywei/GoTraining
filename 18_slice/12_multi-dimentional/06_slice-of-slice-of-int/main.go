package main

import "fmt"

func main() {

	transactions := make([][]int, 0, 3) // slice of slice of int

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4) // [0 1 2 3]
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}
