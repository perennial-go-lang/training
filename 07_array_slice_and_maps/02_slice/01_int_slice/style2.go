package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	for index, value := range x {
		fmt.Println(index, " - ", value)
	}
}
