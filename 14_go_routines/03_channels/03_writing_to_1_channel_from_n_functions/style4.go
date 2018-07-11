package main

import "fmt"

// Up until now, we have seen two functions write to a single channel. In this program, we will
// extrapolate that to `n` functions writing to a single channel.

func main() {
	n := 10 // number of functions to spawn
	c := make(chan int) // unbuffered integer channel
	done := make(chan bool) // semaphore channel

	for i := 0; i < n; i++ {
		go func(n int) {
			for i := 0 + (n*10); i < 10 +(n*10); i++ {
				c <- i
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
