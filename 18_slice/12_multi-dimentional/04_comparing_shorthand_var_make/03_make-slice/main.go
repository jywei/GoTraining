package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Roy"
	// student = append(student, "Roy")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
