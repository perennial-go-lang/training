package main

import "fmt"

func main() {
	var x [30]int

	fmt.Printf("Type: %T\n", x)
	fmt.Println("Contents", x)
	fmt.Println("Length of Array", len(x))
	fmt.Println("Value at 25th position", x[24])
	x[24] = 27
	fmt.Println("Value at 25th position", x[24])
	fmt.Println("Contents", x)
	//x[30] = 100 // Doesn't work. Arrays are of fixed capacity
}
