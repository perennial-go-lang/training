package main

import (
	"fmt"
)

type car struct {
	model   string
	company string
	price   int
}

func (c car) carDescription() string {
	return "Company of car " + c.model + " is " + c.company
}

func (c car) carCompany() string {
	return "Company name is " + c.company
}

func (s showroom) carCompany() string {
	return "Company name is " + s.company
}

type showroom struct {
	car
	company string
}

func main() {
	c1 := car{model: "A", company: "Mercedes", price: 20000}
	fmt.Print(c1.model + " " + c1.company + " ")
	fmt.Println(c1.price)

	fmt.Println(c1.carDescription())

	s1 := showroom{car: car{model: "A8", company: "Audi", price: 5000000}, company: "Lamborghini"}
	fmt.Print(s1.model + " " + s1.car.company + " ")
	fmt.Println(s1.price)

	fmt.Print(s1.model + " " + s1.company + " ")
	fmt.Println(s1.price)

	fmt.Println(c1.carCompany())
	fmt.Println(s1.carCompany())

	c2 := &c1
	fmt.Println("%T %T", c2, c1)
	fmt.Print(c2.model + " " + c2.company + " ")
	fmt.Println(c2.price)

	s2 := &s1
	fmt.Println("%T %T", s2, s1)
	fmt.Print(s2.car.company + " " + s2.company)
}
