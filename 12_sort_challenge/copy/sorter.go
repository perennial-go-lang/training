package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name       string
	Age        int
	Percentage float64
}

type By func(s1, s2 *Student) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(Students []Student) {
	ps := &studentSorter{
		students: Students,
		by:       by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type studentSorter struct {
	students []Student
	by       func(s1, s2 *Student) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *studentSorter) Len() int {
	return len(s.students)
}

// Swap is part of sort.Interface.
func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j])
}

func main() {

	name := func(p1, p2 *Student) bool {
		return p1.Name < p2.Name
	}

	age := func(p1, p2 *Student) bool {
		return p1.Age < p2.Age
	}
	students := []Student{
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
	By(name).Sort(students)
	fmt.Println(students)

	By(age).Sort(students)
	fmt.Println(students)
}
