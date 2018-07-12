package main

import (
	"fmt"
	"program_metrics"
)

func main() {
	var numbersToPrint uint64
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)
	metrics.Start()

	f := fibonacci()

	for i := uint64(0); i <= numbersToPrint; i++ {
		fmt.Print(f(), " ")
	}

	metrics.End()
}

func fibonacci() func() uint64 {
	a, b := uint64(1), uint64(1)
	return func() uint64 {
		defer func() {
			a, b = b, a + b
		}()
		return a
	}
}
