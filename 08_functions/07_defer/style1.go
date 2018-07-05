package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("World")
}
func main() {
	defer hello()
	defer func() {
		fmt.Println("This is anonymous function..")
	}()
	world()
}
