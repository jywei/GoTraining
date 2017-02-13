package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Roy"
	// student = append(student, "Roy")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
