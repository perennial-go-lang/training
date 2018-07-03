package main

import "fmt"

func main() {
	x := 30
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	fmt.Println(x)
}