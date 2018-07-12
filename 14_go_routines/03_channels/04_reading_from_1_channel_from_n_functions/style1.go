package main

import "fmt"

func main() {
	c := make(chan int) //unbuffered integer channel
	done := make(chan bool) // semaphore channel

	go func() {
		//writer function
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		//reader function 1
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		//reader function 2
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	<- done
	<- done
}
