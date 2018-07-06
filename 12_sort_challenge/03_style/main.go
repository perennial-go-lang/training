package main

import (
	. "github.com/perennial-go-training/training/12_sort_challenge/03_style/student"
	"fmt"
	"sort"
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

	fmt.Println("List of students")
	fmt.Println(students)

	fmt.Println("Sorting in Ascending Order")
	sort.Sort(students)
	fmt.Println(students)

	fmt.Println("Sorting in Descending Order")
	sort.Sort(sort.Reverse(students))
	fmt.Println(students)

	fmt.Println("Sorting by Age in Ascending Order")
	sort.Sort(ByAge(students))
	fmt.Println(ByAge(students))

	fmt.Println("Sorting by Age in Descending Order")
	sort.Sort(sort.Reverse(ByAge(students)))
	fmt.Println(ByAge(students))

	fmt.Println("Sorting by Percentage in Ascending Order")
	sort.Sort(ByPercentage(students))
	fmt.Println(ByPercentage(students))

	fmt.Println("Sorting by Percentage in Descending Order")
	sort.Sort(sort.Reverse(ByPercentage(students)))
	fmt.Println(ByPercentage(students))
}
