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

	for i := uint64(0); i <= numbersToPrint; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	metrics.End()
}

func fibonacci(n uint64) uint64 {
	if n <= 1 { return 1 }
	return fibonacci(n - 1) + fibonacci(n - 2)
}