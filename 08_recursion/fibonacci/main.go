package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		aaa := fibonacci(i)
		fmt.Printf("%d ", aaa)
	}
}

func fibonacci(num int) int {
	if num <= 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
