package main

import "fmt"

func main() {
	var number int
	fmt.Printf("Enter number : ")
	fmt.Scanf("%d", &number)
	for i := 1; i <= number; i++ {
		fmt.Println(getFibonacci(i))
	}
}

func getFibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return getFibonacci(number-1) + getFibonacci(number-2)
}
