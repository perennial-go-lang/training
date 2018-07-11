package main

import "fmt"

func main() {
	n := 10
	c := make(chan int) //unbuffered integer channel
	done := make(chan bool) // semaphore channel

	go func() {
		//writer function
		for i := 0; i < 10 * n; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func(func_num int) {
			//reader function 1
			for n := range c {
				fmt.Println("from function", func_num, ":", n)
			}
			done <- true
		}(i)
	}

	for i := 0; i < n; i++ {
		<- done
	}
}
