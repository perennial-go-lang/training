package main

import (
	"fmt"
	"program_metrics"
	"sync"
)
var w sync.WaitGroup

func main() {
	var start, end uint64

	fmt.Println("Enter two numbers")
	fmt.Scan(&start, &end)

	metrics.Start()

	w.Add(int(end) - int(start) + 1)

	for n := range fact(gen(start, end)) {
		fmt.Println(n)
	}

	metrics.End()
}
func gen(start, end uint64) <-chan uint64 {
	out := make(chan uint64)
	go func() {
		for i := start; i <= end; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func fact(in <-chan uint64) <-chan uint64{

	out := make(chan uint64)
	fmt.Println(len(in))

	for i := range in {
		go func(n uint64) {
			var total uint64 = 1
			for j := n; j > 0; j-- {

				total *= j
			}
			out <- total
			w.Done()
		}(i)
	}

	go func() {
		w.Wait()
		close(out)
	}()

	return out
}