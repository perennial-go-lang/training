package main

import "fmt"

func main() {
	//var x [256]int
	var x [256]byte
	//var x [256]string

	fmt.Println("Length of Array", len(x))
	fmt.Println("Value at 24th position", x[25])

	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}

	fmt.Println("Value \t Type \t Binary")
	for i, v := range x {
		fmt.Printf("%v \t\t %T \t %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}