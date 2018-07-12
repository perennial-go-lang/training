package main

import "fmt"

func main() {
	c1 := generator()
	c2 := generator()
	c3 := incrementor(c1)
	c4 := incrementor(c2)

	fmt.Println(<-c3 + <-c4)
}

func generator() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func incrementor(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}