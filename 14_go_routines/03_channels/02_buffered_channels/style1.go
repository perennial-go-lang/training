package main

import "fmt"

func main() {
	c := make(chan int, 10) // buffered channel

	for i := 0; i < 10; i++ {
		c <- i
	}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//}()

	// There is no need to close a channel if you wish to access the elements one by one

	for i := 0; i < 10; i++ {
		fmt.Println(<- c)
	}
}