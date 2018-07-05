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
	By(Name).Sort(students)
	fmt.Println(students)

	fmt.Println("Sorting by Age")
	By(Age).Sort(students)
	fmt.Println(students)

	fmt.Println("Sorting by Percentage")
	By(Percentage).Sort(students)
	fmt.Println(students)
}
