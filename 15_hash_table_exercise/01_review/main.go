package main

import "fmt"

func main() {
	// rune = character = integer
	myLetter := 'A'
	fmt.Println(myLetter)
	fmt.Printf("%T \n", myLetter)

	fmt.Println("B")
	fmt.Printf("%T \n", "B")

	fmt.Println("C"[0])
	fmt.Printf("%T \n", "C"[0])
}
