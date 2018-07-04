package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	fmt.Println("Slice 1: ", slice)
	fmt.Println("Slice:2: ", slice2)

	slice = append(slice, slice2...)

	fmt.Println("Joined Slice: ", slice)

	slice = append(slice[:4], slice[5:]...) // Delete 5

	fmt.Println("After deleting 5 from slice: ", slice1)
}
