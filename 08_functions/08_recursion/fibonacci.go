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

	for i := 0; i <= numbersToPrint; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	metrics.End()
}

func fibonacci(n int) int {
	if n <= 1 { return 1 }
	return fibonacci(n - 1) + fibonacci(n - 2)
}