package main

import (
	"fmt"
	. "github.com/perennial-go-training/training/12_sort_challenge/01_style1/student"
	"sort"
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
	sort.Sort(ByName{students, false})
	fmt.Println(students)
}
