package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	slice3 := append(slice1, slice2...)

	fmt.Printf("Slice after join : ")
	fmt.Println(slice3)

	slice1 = append(slice3[:2], slice3[3:]...)
	fmt.Printf("Slice after deleting : ")
	fmt.Println(slice1)
}
