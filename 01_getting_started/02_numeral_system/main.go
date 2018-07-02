package main

import "fmt"

func main() {
	//fmt.Printf("Decimal - %d \t Binary - %b \t Octal - %o \t Hexadecimal - %#X\n", 42, 42, 42, 42)
	//fmt.Printf("Decimal - %d \t Binary - %b \t Octal - %o \t Hexadecimal - %#x\n", 42, 42, 42, 42)
	//fmt.Printf("Decimal - %d \t Binary - %b \t Octal - %o \t Hexadecimal - %X\n", 42, 42, 42, 42)
	fmt.Printf("Decimal - %d \t Binary - %b \t Octal - %o \t Hexadecimal - %x\n", 42, 42, 42, 42)
	fmt.Printf("Float - %f \t Float accurate to 3 decimal places - %.3f\n", 42.0126232, 42.0126232)
	fmt.Printf("Character (Integer = 65) - %c Integer (Character = 'a')  %d", 65, 'a')
}
