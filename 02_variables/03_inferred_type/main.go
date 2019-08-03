package main

import "fmt"

func main() {
	var hello, world = "Hello", "World!"
	var a, b, c = 1, 2, 3.0 //c is by default initialized as float64
	var d float32 = 2.0

	fmt.Printf("hello(type) - %T (value) - %v\n", hello, hello)
	fmt.Printf("world(type) - %T (value) - %v\n", world, world)
	fmt.Printf("a(type) - %T (value) - %v\n", a, a)
	fmt.Printf("b(type) - %T (value) - %v\n", b, b)
	fmt.Printf("c(type) - %T (value) - %v", c, c)

	//c = c + d //This does not work as c if float64 and d is defined as float32
}
