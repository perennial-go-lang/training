package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter any integer")
	fmt.Scan(&num)
	// fibSeries := fibbonaciSeries(num)
	fmt.Println("Fibonacci Series Val= ", fibbonaciSeries(num))

	for i := 1; i <= num; i++ {
		fmt.Println(fibbonaciSeries(i))
	}

}

func fibbonaciSeries(n int) int {
	if n <= 1 {
		return n
	}
	return fibbonaciSeries(n-1) + fibbonaciSeries(n-2)
}
