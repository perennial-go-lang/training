package main

import "fmt"

func main() {
	c := make(chan int) // unbuffered channel

	// An unbuffered channel always needs two threads of execution. One to push, one to pull. In absence of
	// either, there is a deadlock

	//for i := 0; i < 10; i++ {
	//	c <- i
	//}

	// In this scenario, there is go routine (sub thread) that pushes data on to the channel and the main
	// thread that pulls data off the channel. No deadlock.

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// There is no need to close a channel if you wish to access the elements one by one

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}