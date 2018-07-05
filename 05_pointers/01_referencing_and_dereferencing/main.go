package main

import "fmt"

func main() {
	a := 24 // Declare an integer variable with value 24
	b := &a // Declare an integer pointer that references the address of a

	fmt.Println(a, *b) // Print 24 24

	c := *b // Declare variable c that dereferences the value from the integer pointer b
	a = 30  // Update value of variable a

	fmt.Println(a, *b, c) // What will this print?
}
