package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4}
	y := []int{5, 6, 7, 8, 9}

	x = append(x, y...)

	fmt.Println(x)

	x = append(x[:5], x[6:]...) // Delete 5

	fmt.Println(x)
}
