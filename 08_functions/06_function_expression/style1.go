package main

import "fmt"

func main() {
	greet := func() {
		fmt.Println("Hello, World!")
	}

	greet()
	fmt.Printf("%T", greet)
}
