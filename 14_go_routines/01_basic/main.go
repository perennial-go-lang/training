package main

import (
	"fmt"
)

func main() {
	message := make(chan string)

	go func() { message <- "Hiiiii" }()

	msg := <-message
	fmt.Println(msg)
}
