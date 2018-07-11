package main

import "fmt"

func main() {
	average(1, 2, 3, 4, 5)

	numbers := []int{6, 7, 8, 9, 10}
	average(numbers...)
	sum(numbers)
}

func average(num ...int) {
	var avg int
	for _, v := range num {
		avg += v
	}
	avg = avg / len(num)
	fmt.Println(avg)
}

func sum(num []int) {
	var sum int
	for _, v := range num {
		sum += v
	}
	fmt.Println(sum)
}
