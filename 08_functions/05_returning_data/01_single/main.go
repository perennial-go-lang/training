package main

import "fmt"

func main() {
	x := 2
	x = square(x)
	y := cube(x)
	fmt.Println(x, y)
}

func square(x int) int {
	return x * x
}

// named return
func cube(x int) (c int) {
	c = x * x * x
	return
}
