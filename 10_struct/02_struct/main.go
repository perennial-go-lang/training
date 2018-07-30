package main

import (
	"fmt"
)

type person struct {
	fName string //fields
	lName string
	age   int
}

type student struct {
	person // promotion of person
	school string
	fName  string
}

func (p person) fullName() string {
	return p.fName + " " + p.lName
}

func (p person) greet(greeting string) string {
	return greeting + " " + p.fName + " " + p.lName
}

func (p person) introduce() string {
	return "My name is " + p.fullName()
}

func (s student) introduce() string {
	//return s.person.introduce() + ". I go to " + s.school
	return "My name is " + s.fName + ". I go to " + s.school
}

func main() {
	p1 := person{"A", "B", 22}
	p2 := person{"X", "Y", 33}

	fmt.Println(p1.fName, p1.lName, p1.age)
	fmt.Println(p1.fullName())
	fmt.Println(p1.greet("Good Morning"))

	fmt.Println(p2.fName, p2.lName, p2.age)
	fmt.Println(p2.fullName())
	fmt.Println(p2.greet("Hello"))

	s1 := student{person{"Babulal", "Choudhary", 26}, "", "NewName"}
	s1.school = "ads"
	fmt.Println(s1.fullName())
	fmt.Println(s1.greet("Good Morning"))
	fmt.Println(s1.school)

	fmt.Println(p1.introduce())
	fmt.Println(s1.introduce(), s1.person.introduce())

	fmt.Println(s1.fName, s1.person.fName, s1.lName, s1.person.lName)

	p3 := &p1

	fmt.Println(p3, p1)
	fmt.Printf("%T %T", p3, p1)

	fmt.Println(p3.fName, p1.fName)
}
