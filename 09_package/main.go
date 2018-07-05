package main

import (
	"fmt"

	"github.com/rohitpavaskar/training/09_package/numutil"
	"github.com/rohitpavaskar/training/09_package/stringutil"
)

func main() {
	var source string
	var number int
	fmt.Printf("Enter string to Reverse: ")
	fmt.Scan(&source)
	fmt.Printf("Reversed String: " + stringutil.ReverseString(source) + "\n")
	fmt.Printf("Enter number to check Odd/Even : ")
	fmt.Scan(&number)
	fmt.Printf("Number is:" + numutil.OddOrEven(number) + "\n")
}
