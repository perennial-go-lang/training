package main

import "fmt"

func main() {
	x := make([]int, 0, 10)

	//x[1] = 1 // Panic! Use append instead.

	fmt.Printf("Type - %T\n", x)
	fmt.Println("Contents - ", x)
	fmt.Println("Length - ", len(x))
	fmt.Println("Capacity - ", cap(x))

	for i := 0; i < 100; i++ {
		x = append(x, i)
		fmt.Printf("Length - %v | Capacity - %v | Value - %v\n", len(x), cap(x), x[i])
	}
}
