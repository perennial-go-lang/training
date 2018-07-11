package main

import (
	"time"
	"fmt"
)

func main() {
	c := make(chan int) // make an unbuffered integer channel

	// writer function 1
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// writer function 2
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// We are trying to access a channel from main thread. This should work like in #1 but does not
	// because now we have multiple processes writing and main trying to read before either of them
	// is finished. We now have two options:
	//
	// 1. Setup another go routine to retrieve the data and print it out
	// 2. Close the channels and access from main
	//for {
	//	fmt.Println(<-c)
	//}

	go func() {
		fmt.Println(<-c)
	}()

	time.Sleep(1 * time.Second)
}

