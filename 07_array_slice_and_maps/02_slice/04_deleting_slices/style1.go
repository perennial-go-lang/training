package main

import (
	"fmt"
)

func main() {
	slice := []string{ "January", "February", "March", "April", "May", "June", "July" }
	slice2 := []string{ "August", "September", "October", "November", "December" }

	slice = append(slice, slice2...)

	fmt.Println(slice)

	slice = append(slice[:2], slice[3:]...) // Delete March

	fmt.Println(slice)
}
