//Factorial using channels and Pipeline and Workgroup
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
	fmt.Println("Factorials Below:")
	var startTime = time.Now()

	//c1 := generator(start, end)
	//c2 := factorial(c1)
	c2 := factorial(generator(start, end), end)
	for n := range c2 {
		fmt.Println(n)
	}
	fmt.Println("Execution Time: ", time.Since(startTime))
}

func factorial(fc <-chan int, end int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(end)

	for num := range fc {
		go func(n int) {
			total := 1
			for i := n; i > 0; i-- {
				total = total * i
			}
			out <- total
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func generator(s, e int) <-chan int {
	out := make(chan int)
	go func() {
		var i int
		for i = s; i <= e; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
