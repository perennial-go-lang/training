package main

import "fmt"

func main() {
	fmt.Println(num)
	fmt.Println(other)
	num := 30 //must be declared before it is called
}

var other = 24
