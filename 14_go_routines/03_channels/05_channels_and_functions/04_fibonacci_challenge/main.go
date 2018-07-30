package main

import (
	"fmt"
)

func main() {
	var numbersToPrint uint64
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)
	// metrics.Start()

	series := fibonacci_semaphore(numbersToPrint)

	for n := range series {
		fmt.Print(n, " ")
	}

	// metrics.End()
}

func fibonacci_semaphore(numbers uint64) <-chan uint64 {
	c := make(chan uint64)

	go func(num uint64) {
		// out := make(chan uint64)
		done := make(chan bool)
		for i := uint64(0); i <= num; i++ {
			go func(n uint64) {
				var t = uint64(1)
				for i, j, k := uint64(1), uint64(1), uint64(0); k < n; i, j, k = i+j, i, k+1 {
					t = i
				}
				c <- t
				done <- true
			}(i)
		}
		go func() {
			for i := uint64(0); i <= num; i++ {
				<-done
			}
			close(c)
		}()
		// for n := range out {
		// 	c <- n
		// }
		// close(c)
	}(numbers)

	return c
}
