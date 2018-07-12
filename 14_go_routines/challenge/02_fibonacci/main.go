package main

import (
	"fmt"
	"time"
)

func main() {
	var numbersToPrint int
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)

	// f := fibonacci()
	var startTime = time.Now()

	fibChan := make(chan int)
	go func() {
		for {
			fibonacci(fibChan)
		}
	}()

	for i := 0; i <= numbersToPrint; i++ {
		fmt.Print(<-fibChan, " ")
	}
	fmt.Println("\nTime Consumed:", time.Since(startTime))
}

func fibonacci(fibChan chan<- int) func() int {
	a, b := 1, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		fibChan <- a
		return a
	}
}
