package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// Deadlock
// 1. Define c as a buffered channel
// 2. wrap c <- 1 in a go routine