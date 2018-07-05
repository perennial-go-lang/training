package main

import (
	"fmt"
	"math"
)

func main() {

	filterMap := make(map[string]func(n int) bool)

	filterMap["DivisbleByTwo"] = func(n int) bool {
		return (n%2 == 0)
	}
	filterMap["OddNumber"] = func(n int) bool {
		return (n&1 == 1)
	}
	filterMap["EvenNumber"] = func(n int) bool {
		return (n&1 == 0)
	}
	filterMap["PrimeNumber"] = func(n int) bool {
		for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
			if n%i == 0 {
				return false
			}
		}
		return n > 1
	}

	numbers := []int{2, 5, 6, 7, 8, 9, 11, 12, 15, 18, 21, 22, 25, 30, 35, 40}

	fmt.Println("DivisiableByTwo : ", filterNumbers(numbers, filterMap["DivisbleByTwo"]))
	fmt.Println("Odd Numbers :", filterNumbers(numbers, filterMap["OddNumber"]))
	fmt.Println("Even Numbers :", filterNumbers(numbers, filterMap["EvenNumber"]))
	fmt.Println("Prime Numbers :", filterNumbers(numbers, filterMap["PrimeNumber"]))
}
func filterNumbers(num []int, condition func(n int) bool) []int {

	filterNumber := make([]int, 0)
	for _, v := range num {
		if condition(v) {
			filterNumber = append(filterNumber, v)
		}
	}
	return filterNumber
}
