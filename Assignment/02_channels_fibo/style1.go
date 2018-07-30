package main

import (
	"fmt"
	"time"
)

func main() {
	var numbersToPrint uint64
	fmt.Println("Enter numbers to print in the Fibonacci Series")
	fmt.Scan(&numbersToPrint)

	var startTime = time.Now()

	out := generateFibo(numbersToPrint)
	for n := range out {
		fmt.Print(n, " ")
	}

	fmt.Println("\nTime consumed: ", time.Since(startTime))

}
func generateFibo(numbers uint64) <-chan uint64 {

	out := make(chan uint64)
	go func(num uint64) {
		// ic := make(chan uint64)
		done := make(chan bool)
		for i := uint64(0); i <= num; i++ {
			go func(n uint64) {
				var t = uint64(1)
				for i, j, k := uint64(1), uint64(1), uint64(0); k < n; i, j, k = i+j, i, k+1 {
					t = i
				}
				out <- t
				done <- true
			}(i)
		}

		go func() {
			for i := uint64(0); i <= num; i++ {
				<-done
			}
			close(out)
		}()
		// for n := range ic {
		// 	out <- n
		// }
		// close(out)
	}(numbers)
	return out
}
