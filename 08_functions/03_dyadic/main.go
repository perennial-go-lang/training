package main

import "fmt"

func main() {
	add(3,5)
	multiply(9,7)
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func multiply(a, b int) {
	fmt.Println(a * b)
}