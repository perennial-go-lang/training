package main

import (
	"fmt"
	"program_metrics"
)

func main() {
	var start, end uint64
	fmt.Println("Enter two numbers")
	fmt.Scan(&start, &end)
	metrics.Start()
	for i := start; i <= end; i++ {
		c := factorial(i)
		for n := range c {
			fmt.Println(n)
		}
	}
	metrics.End()
}

func factorial(n uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		var total uint64 = 1
		for i := n; i > 0; i-- {
			total = total * i
		}
		out <- total
		close(out)
	}()
	return out
}