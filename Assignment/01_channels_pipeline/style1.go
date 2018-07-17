package main

import (
	"fmt"
	"time"
)

func main() {
	var start, end int
	fmt.Println("Enter two numbers")
	fmt.Scan(&start, &end)

	var startTime = time.Now()

	c := factorial(multiFact(start, end))

	for n := range c {
		fmt.Println(n)
	}
	fmt.Println("\nTime consumed: ", time.Since(startTime))

}

func multiFact(start, end int) chan int {
	out := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			total := 1
			for i := num; i > 0; i-- {
				total = total * i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
