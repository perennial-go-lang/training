package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var start, end int
	fmt.Println("Enter two numbers")
	fmt.Scan(&start, &end)

	st := time.Now()

	finalR := factorial(generator(start, end), end)
	for n := range finalR {
		fmt.Println(n)
	}
	fmt.Println(time.Since(st))

}

func factorial(in chan int, end int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(end)
	for n := range in {
		go func(np int) {
			total := 1
			for i := np; i > 0; i-- {
				total = total * i
			}
			out <- total
			wg.Done()
		}(n)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func generator(start, end int) chan int {
	out := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
