package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum, average := sumAndAverage(numbers...)
	fmt.Println(sum, average)
}

func sumAndAverage(num... int) (sum int, avg int) {
	for _, v := range num {
		sum += v
	}
	avg = sum / len(num)
	return
}
