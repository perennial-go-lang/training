package main

import (
	"fmt"
	. "github.com/perennial-go-training/training/12_sort_challenge/01_style1/student"
)

func main() {
	students := Students{
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

	By(Name).SortAsc(students)
	fmt.Println("Ascending Order", students)

	By(Name).SortDesc(students)
	fmt.Println("Descending Order", students)

	fmt.Println("Sorting by Age")

	By(Age).SortAsc(students)
	fmt.Println("Ascending Order", students)

	By(Age).SortDesc(students)
	fmt.Println("Descending Order", students)

	fmt.Println("Sorting by Percentage")

	By(Percentage).SortAsc(students)
	fmt.Println("Ascending Order", students)

	By(Percentage).SortDesc(students)
	fmt.Println("Descending Order", students)
}
