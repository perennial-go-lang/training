package main

import (
	"fmt"
	"sort"

	. "github.com/rohitpavaskar/training/12_sort_challenge/students"
)

func main() {
	students := Students{
		{
			"P",
			20,
			50,
		},
		{
			"BD",
			19,
			60,
		},
		{
			"BA",
			19,
			60,
		},
		{
			"C",
			16,
			40,
		},
	}

	fmt.Println("Student's List")
	fmt.Println(students)
	fmt.Println("")

	fmt.Println("Student's List By Name Ascending")
	sort.Sort(ByName(students))
	fmt.Println(students)
	fmt.Println("Student's List By Name Descending")
	sort.Sort(sort.Reverse(ByName(students)))
	fmt.Println(students)
	fmt.Println("")

	fmt.Println("Student's List By Age Ascending")
	sort.Sort(ByAge(students))
	fmt.Println(students)
	fmt.Println("Student's List By Age Descending")
	sort.Sort(sort.Reverse(ByAge(students)))
	fmt.Println(students)
	fmt.Println("")

	fmt.Println("Student's List By Percentage Ascending")
	sort.Sort(ByPercentage(students))
	fmt.Println(students)
	fmt.Println("Student's List By Percentage Descending")
	sort.Sort(sort.Reverse(ByPercentage(students)))
	fmt.Println(students)
}
