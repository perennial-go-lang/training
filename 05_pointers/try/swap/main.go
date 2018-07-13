package main

import (
	"fmt"
)

func main() {
	x := 2
	y := 3
	swap(&x, &y)
	fmt.Println(x, y)

	p := 1.5
	square(&p)
	fmt.Println(p)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
func square(x *float64) {
	*x = *x * *x
}
