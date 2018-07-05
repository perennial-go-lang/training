package main

import "fmt"

type number int

func main() {
	var a number
	a = 30

	fmt.Printf("%T %v", a, a)

	var b int
	b = 24
	fmt.Printf("%T %v", b, b)

	fmt.Println(a + number(b))
}