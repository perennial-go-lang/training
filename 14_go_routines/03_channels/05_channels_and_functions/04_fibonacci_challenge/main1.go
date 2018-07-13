package main

import (
	"fmt"
	"program_metrics"
	"sync"
)

func main() {
	var numbersToPrint uint64
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)
	metrics.Start()

	series := fibonacci(numbersToPrint)

	for n := range series {
		fmt.Print(n, " ")
	}

	metrics.End()
}

func fibonacci(numbers uint64) <-chan uint64 {
	c := make(chan uint64)

	go func(num uint64) {
		var wg sync.WaitGroup

		wg.Add(int(num)+1)
		for i := uint64(0); i <= num; i++ {
			go func(n uint64) {
				var t = uint64(1)
				for i, j, k := uint64(1), uint64(1), uint64(0); k < n ; i, j, k = i+j,i, k+1 {
					t = i
				}
				c <- t
				wg.Done()
			}(i)
		}
		go func() {
			wg.Wait()
			close(c)
		}()
	}(numbers)

	return c
}

