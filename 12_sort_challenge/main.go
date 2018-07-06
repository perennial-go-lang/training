package main

import (
	"fmt"
	"sort"
)

type Person string

type People []Person

func main() {
	var person Person = "Ajitem"
	var people = People{"Ankit", "Soni", "Rohit"}

	people = append(people, person)
	person.who()
	fmt.Println(people)

	sort.Sort(people)

	fmt.Println(people)

	sort.Sort(sort.Reverse(people))

	fmt.Println(people)
	fmt.Printf("%T %T", people, sort.Reverse(people))
}

func (p People) Len() int {
	return  len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Person) who() {
	fmt.Println("I am", p)
}