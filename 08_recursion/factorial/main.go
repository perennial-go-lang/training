package main

import "fmt"

func main() {
	numbers := []int{4, 2, 3, 4, 5}

	// OddNumbers := func(num int) bool {
	// 	return (num & 1) == 1
	// }
	EvenNumbers := func(num int) bool {
		return (num & 1) == 0
	}
	fmt.Println(basicOperations(numbers, OddNumber))
	fmt.Println(basicOperations(numbers, EvenNumbers))
	// fmt.Println(OddNumber(5))
}

func OddNumber(num int) bool {
	return (num & 1) == 1
}

func basicOperations(numbers []int, operation func(number int) bool) (result []int) {

	for _, num := range numbers {
		if operation(num) {
			result = append(result, num)
		}
	}
	return
}
