package main

import (
	"fmt"

	nume "github.com/soniraju/training/09_package/numutil"
	strg "github.com/soniraju/training/09_package/stringutil"
)

func main() {
	var str string
	var num int

	fmt.Printf("Enter String To Reverse: ")
	fmt.Scan(&str)
	fmt.Println("Reversed String:= ", strg.ReverseString(str))

	fmt.Printf("Enter Integer Number: ")
	fmt.Scan(&num)
	fmt.Println("Given Number is ", nume.OddOrEven(num))
}
