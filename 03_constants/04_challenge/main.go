package main

import "fmt"

const (
	kb = (iota + 2) * 1024
	mb
	gb
	tb
)

func main() {
	fmt.Println(kb)

	// fmt.Print(mb)
	// fmt.Print(" | ")
	// fmt.Println(kb * 1024)

	// fmt.Print(gb)
	// fmt.Print(" | ")
	// fmt.Print(mb * 1024)
	// fmt.Print(" | ")
	// fmt.Println(kb * 1024)

	// fmt.Println(gb)
	// fmt.Println(tb)
}
