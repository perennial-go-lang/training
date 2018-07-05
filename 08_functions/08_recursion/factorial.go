package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * factorial(n - 1)
}
