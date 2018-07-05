package main

import "fmt"

func main() {
	for i := 1; i <= 10; i = i + 1 {
		k := fibo(i)
		fmt.Print(k)
	}
}

func fibo(n int) int {
	if n <= 1 {
		return 1
	}
	return (fibo(n-1) + fibo(n-2))
}
