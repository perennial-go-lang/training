package main

import "fmt"

func main() {
	var cars [][]string

	car1 := make([]string, 3)
	car1[0] = "Ford"
	car1[1] = "New"
	car1[2] = "2,00,000"

	cars = append(cars, car1)

	car2 := make([]string, 3)
	car2[0] = "Maruti"
	car2[1] = "Used"
	car2[2] = "50,000"

	cars = append(cars, car2)

	fmt.Println(cars)
}
