package main

import "fmt"

func main() {
	var numbersToPrint int
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)
	for i := 0; i <= numbersToPrint; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

func fibonacci(n int) int {
	if n <= 1 { return 1 }
	return fibonacci(n - 1) + fibonacci(n - 2)
}