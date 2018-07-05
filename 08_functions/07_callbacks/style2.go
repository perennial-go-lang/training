package main

import "fmt"

func main() {
	numbers := []int{123,2,54,21,53,4,22,13,354,1,6,32,5}

	criteria := make(map[string] func (n int) (bool))

	criteria["isDivisibleByThree"] = func(n int) bool {
		return n % 3 == 0
	}

	divByThree := filterNumbers(numbers, criteria["isDivisibleByThree"])

	oddNumbers := filterNumbers(numbers, func(n int) bool {
		return n & 1 == 1
	})

	fmt.Println(divByThree)
	fmt.Println(oddNumbers)
}

func filterNumbers(numbers []int, criteria func(n int) bool) []int {
	validNumbers := make([]int, 0)
	for _, n := range numbers {
		if valid := criteria(n); valid {
			validNumbers = append(validNumbers, n)
		}
	}
	return validNumbers
}
