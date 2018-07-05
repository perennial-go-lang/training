package main

import "fmt"

func main() {
	var hello, world, message string
	var a, b, c int
	var d float64 = 4.99

	hello = "Hello,"
	world = "World!"

	a = 1

	fmt.Printf("hello(type) - %T (value) - %v\n", hello, hello)
	fmt.Printf("world(type) - %T (value) - %v\n", world, world)
	fmt.Printf("message(type) - %T (value) - %v\n", message, message)
	fmt.Printf("a(type) - %T (value) - %v\n", a, a)
	fmt.Printf("b(type) - %T (value) - %v\n", b, b)
	fmt.Printf("c(type) - %T (value) - %v", c, c)
	fmt.Printf("d(type) - %T (value) - %v", d, d)
}
