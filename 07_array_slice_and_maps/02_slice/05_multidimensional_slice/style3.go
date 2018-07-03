package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)

	fmt.Println(student)
	fmt.Println(students)

	student[0] = "Ajitem"

	fmt.Println(student)
	fmt.Println(students)

	fmt.Println(student == nil)
}