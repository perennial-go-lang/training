package main

import "fmt"

var num = 30

func main() {
	fmt.Println(num)
	foo()
	sum()
}

func foo() {
	num++
	fmt.Println(num)
}
