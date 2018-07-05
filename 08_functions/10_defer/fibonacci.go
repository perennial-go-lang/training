package main

import "fmt"

func main() {
	var numbersToPrint int
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)

	f := fibonacci()

	for i := 0; i <= 10; i++ {
		//_ = f()
		fmt.Print(f(), " ")
	}
	//fmt.Print(f())
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
