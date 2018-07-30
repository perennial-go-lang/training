package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	done := make(chan bool)
	go worker(done)
	fmt.Scanf("", &name)
	fmt.Println("")
	<-done
}
func worker(done chan bool) {
	fmt.Println("Working.....")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- true
}
