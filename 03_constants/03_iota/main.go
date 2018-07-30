package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	_ = iota
	x
	y
	z
	p
	q
	r
	s
)

 func main() {
 	fmt.Println(a, b, c)
 	fmt.Println( x, y, z, p, q, r, s)
 }
