package main

import (
	"fmt"
)

type vehicle struct {
	manufacturer string
	wheels       int
	releaseDate  string
}

type car struct {
	vehicle
	model       string
	releaseDate string
}

func (v vehicle) getmanufacturer() string {
	return v.manufacturer
}

func (c car) getModel() string {
	return c.model
}

func (c car) getVehicleReleaseDate() string {
	return c.vehicle.releaseDate
}

func main() {
	v1 := vehicle{"Tata", 4, "1960"}
	fmt.Println(v1.manufacturer, v1.wheels)
	c1 := car{v1, "Nano", "2012"}
	c2 := car{vehicle{manufacturer: "Ford", wheels: 4, releaseDate: "1850"}, "Nano", "2012"}
	fmt.Println(c2)
	fmt.Println(c1)

	fmt.Println(c1.getmanufacturer())
	fmt.Println(c1.getModel())
	fmt.Println(c1.getVehicleReleaseDate())
}
