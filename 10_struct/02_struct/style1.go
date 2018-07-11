package main

import (
	"fmt"
)

type person struct {
	fName string //fields
	lName string
	age   int
}

//inheritance
type employee struct {
	person      //Promotion of employee
	designation string
}

func (p person) fullName() string {
	return p.fName + " " + p.lName
}
func (p person) introduce() string {
	return "My Name is " + p.fullName()
}

func (e employee) introduce() string {
	return "My Name is " + e.fullName() + ". I am works as " + e.designation
}

func main() {
	p1 := person{fName: "Babulal", lName: "Choudhary", age: 26}
	e1 := employee{person: person{lName: "Khade", fName: "Akshay", age: 25}, designation: "Software Engineer"}

	fmt.Println(p1.fullName())
	fmt.Println(e1.fullName())

	//Method overloading
	fmt.Println(p1.introduce())
	fmt.Println(e1.introduce())

	//Pointer
	e2 := &e1
	fmt.Printf("Type : %T %v \n", e1, e1)
	fmt.Printf("Type : %T %v \n", e2, e2)
}
