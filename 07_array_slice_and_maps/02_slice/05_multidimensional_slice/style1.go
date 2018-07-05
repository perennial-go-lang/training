package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}

	fmt.Println(student)
	fmt.Println(students)

	fmt.Println(student == nil)

	//student[0] = "Ajitem"
	student = append(student, "Ajitem")

	fmt.Println(student)
	fmt.Println(students)

}