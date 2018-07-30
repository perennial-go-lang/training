package main

import (
	"fmt"
)

func main() {
	message := make(chan string, 2)

	message <- "HI"
	message <- "Rohit"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
