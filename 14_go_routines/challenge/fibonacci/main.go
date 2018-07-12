package main

import (
	"fmt"
	"time"
)

func main() {
	var numbersToPrint int
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)

	f := fibonacci()
	sem := make(chan bool)
	var ss = time.Now()
	for i := 0; i <= numbersToPrint; i++ {
		go func() {
			fmt.Print(f(), " ")
			sem <- true
		}()
	}

	for i := 0; i <= numbersToPrint; i++ {
		<-sem
	}
	fmt.Println("\nTime Consumed:", time.Since(ss))
}

func fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}
