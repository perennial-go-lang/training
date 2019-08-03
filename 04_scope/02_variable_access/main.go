package main

import "fmt"

func main() {
	x := 30
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x //as x wasn't declared in current scope
	// this does not compile
	fmt.Println(x)
}
