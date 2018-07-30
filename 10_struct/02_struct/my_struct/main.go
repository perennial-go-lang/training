package main

import (
	"fmt"
	"strconv"
)

type person struct {
	fName, lName, compName string
}

type employee struct {
	person //Promoting person struct in employee struct
	empId  int
	fName  string
}

func (p person) getDetails() string {
	return "Name: " + p.fName + " " + p.lName + "\nComapny: " + p.compName
}

func (e employee) getDetails() string {
	return "Emp Id:" + strconv.Itoa(e.empId) + "\nName: " + e.fName + " " + e.lName + "\nComapny: " + e.compName
}

func main() {
	p1 := person{fName: "Soni", lName: "Raju", compName: "Perennial"}
	fmt.Println(p1.getDetails())

	fmt.Println("")

	e1 := employee{person: person{fName: "SSS", lName: "Gandhi", compName: "Perennial System"}, empId: 120, fName: "Sonii"}
	fmt.Println(e1.getDetails())
}
