package main

import (
	"fmt"
	"program_metrics"
)

func main() {
	var numbersToPrint int
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)
	metrics.Start()

	f := fibonacci()

	for i := 0; i <= 10; i++ {
		fmt.Print(f(), " ")
	}

	metrics.End()
}

func fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		defer func() {
			a, b = b, a + b
		}()
		return a
	}
}
