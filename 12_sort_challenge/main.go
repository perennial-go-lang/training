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

// ByAge implements sort.Interface based on the Age field.
type ByAge []Student

type ByName []Student
type ByPercentage []Student

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByPercentage) Len() int           { return len(a) }
func (a ByPercentage) Less(i, j int) bool { return a[i].Percentage > a[j].Percentage }
func (a ByPercentage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	students := []Student{
		{
			"Akshay",
			20,
			50,
		},
		{
			"Rohit",
			19,
			60,
		},
		{
			"Babulal",
			26,
			68,
		},
	}
	sort.Sort(ByAge(students))
	fmt.Println(students)

	sort.Sort(ByName(students))
	fmt.Println(students)

	sort.Sort(ByPercentage(students))
	fmt.Println(students)
}
