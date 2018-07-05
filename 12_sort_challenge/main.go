package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age int
	Percentage float64
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {

}

func (s Students) Less(i, j int) bool {
	return true
}

func main() {
	students := []Student{
		{
			"A",
			20,
			50,
		},
		{
			"B",
			19,
			60,
		},
	}
	sort.Sort(Students(students))
	fmt.Println(students)
}