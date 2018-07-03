package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	x = iota
	y
	z
)

 func main() {
 	fmt.Println(a, b, c)
 	fmt.Println(x, y, z)
 }
