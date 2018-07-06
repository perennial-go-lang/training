package main

import (
	"fmt"

	. "github.com/rohitpavaskar/training/12_sort_challenge/multisort/student"
)

func main() {
	fmt.Println("")

	students := Students{
		Student{"BD", 26, 10.5},
		Student{"BA", 22, 50.0},
		Student{"A", 28, 30.5},
		Student{"A", 11, 30.5},
		Student{"A", 40, 9.5},
	}

	name := func(s1, s2 Student) bool {
		return s1.Name < s2.Name
	}

	age := func(s1, s2 Student) bool {
		return s1.Age < s2.Age
	}

	percentage := func(s1, s2 Student) bool {
		return s1.Percentage < s2.Percentage
	}

	fmt.Println("Student's List")
	fmt.Println(students)
	fmt.Println()

	fmt.Println("Student's List by name ascending")
	By{name}.SortAsc(students)
	fmt.Println(students)

	fmt.Println("Student's List by name, percentage, age ascending")
	By{name, percentage, age}.SortAsc(students)
	fmt.Println(students)

	fmt.Println("Student's List by age, percentage, name ascending")
	By{age, percentage, name}.SortAsc(students)
	fmt.Println(students)
}
