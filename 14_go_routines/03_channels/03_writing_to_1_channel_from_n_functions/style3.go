package main

import (
	"fmt"
)

// SEMAPHORE - Semaphore is a variable that is used by a program / OS to manage access to a system
// resource

func main() {
	c := make(chan int) // make an unbuffered integer channel
	done := make(chan bool) // semaphore channel

	// writer function 1
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// writer function 2
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}