package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var average float64 = 0.0
	for i := 0; i < len(numbers); i++ {
		average += float64(numbers[i])
	}
	fmt.Println(`average is
	`, average)
}
