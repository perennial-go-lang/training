package main

import "fmt"

func main() {
	fmt.Printf("Float: %f, Decimal: %d, Binary: %b, Octal: %o\n",
		30.23, 30, 30, 30)
	fmt.Printf("Hexadecimal(Capital): %X ,Hexadecimal(small): %x ,Hexadecimal(0x): %#x\n", 30, 30, 30)
	fmt.Printf("Characher: %c, Character(Integer): %c\n", 'A', 65)
}
