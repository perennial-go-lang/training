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
	p
	q
	r
	s
)

 func main() {
 	fmt.Println(a, b, c)
 	fmt.Println( y, z, p, q, r, s)
 }
