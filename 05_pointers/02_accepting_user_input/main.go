package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Enter two numbers:")
	fmt.Scan(&num1, &num2)

	fmt.Println("Sum of the two numbers is: ", num1 + num2)
}