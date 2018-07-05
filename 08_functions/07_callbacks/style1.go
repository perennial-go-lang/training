package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4,5}

	readAndPrint(numbers, func(n int) {
		fmt.Println(n)
	})
}

func readAndPrint(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
