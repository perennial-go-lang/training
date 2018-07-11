package main

import (
	"fmt"

	. "github.com/rohitpavaskar/training/12_sort_challenge/main/student"
)

func main() {
	fmt.Println("")

	students := Students{
		Student{"BD", 26},
		Student{"BA", 22},
		Student{"CD", 28},
		Student{"A", 40},
	}

	name := func(s1, s2 Student) bool {
		return s1.Name < s2.Name
	}

	age := func(s1, s2 Student) bool {
		return s1.Age < s2.Age
	}

	fmt.Println("Student's List")
	fmt.Println(students)
	fmt.Println()

	fmt.Println("Student's List by name ascending")
	By(name).SortAsc(students)
	fmt.Println(students)
	fmt.Println("Student's List by name descending")
	By(name).SortDesc(students)
	fmt.Println(students)
	fmt.Println()

	fmt.Println("Student's List by age ascending")
	By(age).SortAsc(students)
	fmt.Println(students)
	fmt.Println("Student's List by age descending")
	By(age).SortDesc(students)
	fmt.Println(students)
	fmt.Println()
}
