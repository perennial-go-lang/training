package main

import (
	"fmt"
	. "github.com/perennial-go-training/training/12_sort_challenge/02_style2/student"
)

func main() {
	students := []Student{
		{
			"B",
			20,
			50,
		},
		{
			"A",
			19,
			60,
		},
		{
			"C",
			18,
			55,
		},
	}
	fmt.Println("Students List")
	fmt.Println(students)

	fmt.Println("Sorting by Name")

	By(Age).Sort(students)
	fmt.Println("Ascending Order", students)
	By(Age, Percentage).Sort(students)
	fmt.Println("Ascending Order", students)

	//By(Name).Sort(student, "desc")
	//fmt.Println("Descending Order", student)
}
