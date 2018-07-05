package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x)

	fmt.Println(x[2:4])  // slicing a slice
	fmt.Println(x[:4])  // slicing a slice
	fmt.Println(x[2:])  // slicing a slice
	fmt.Println(x[2])    // index access; accessing by index

	fmt.Println("myString"[2:]) // index access; accessing by index
}