package main

import "fmt"

func main() {
	c := make(chan int, 10) // buffered channel

	for i := 0; i < 10; i++ {
		c <- i
	}
	
	// The channel must be properly closed in order to access it in a loop with range.
	close(c)

	// In this scenario, there is go routine (sub thread) that pushes data on to the channel and the main
	// thread that pulls data off the channel. No deadlock.

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//	// The channel must be properly closed in order to access it in a loop with range.
	//	close(c)
	//}()


	for n := range c {
		fmt.Println(n)
	}
}

// Fatal Error