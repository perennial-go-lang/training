package main

import "fmt"

func main() {
	message := "Hello, World!"
	a, b, c := 1, 22.7, true

	fmt.Printf("message(type) - %T (value) - %v\n", message, message)
	fmt.Printf("a(type) - %T (value) - %v\n", a, a)
	fmt.Printf("b(type) - %T (value) - %v\n", b, b)
	fmt.Printf("c(type) - %T (value) - %v", c, c)
}
