package main

import "fmt"

func main() {
	var student []string // var will always set it to zero value
	var students [][]string
	student[0] = "Roy"
	// student = append(student, "Roy")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
