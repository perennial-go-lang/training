package main

import (
	. "github.com/perennial-go-training/training/11_interface/shape"
	"fmt"
	"github.com/perennial-go-training/training/11_interface/figures"
)

func main() {
	var shape Shape

	fmt.Printf("%T\n", shape)

	shape = figures.Square{5}
	fmt.Printf("%T\n", shape)
	fmt.Println(printArea(shape))

	shape = figures.Rectanagle{5, 6}
	fmt.Printf("%T\n", shape)
	fmt.Println(printArea(shape))

	shape = figures.Circle{10}
	fmt.Printf("%T\n", shape)
	fmt.Println(printArea(shape))

	shape = triangle{10, 15}
	fmt.Printf("%T\n", shape)
	fmt.Println(printArea(shape))
}

func printArea(shape Shape) float64{
	return shape.Area()
}
