package main

import (
	"fmt"
	strrev "github.com/babulal107/training/09_package/stringutil"
	nutil "github.com/babulal107/training/09_package/numutil"
)

func main() {
	var name string
	var num int

	fmt.Printf("Enter string : ")
	fmt.Scan(&name)
	reversedStr := strrev.ReverseString(name)
	fmt.Printf("Reversed string : %s \n", reversedStr)

	fmt.Printf("Enter Number : ")
	fmt.Scan(&num)
	fmt.Printf("Number (%d) is : %s \n", num, nutil.OddOrEven(num))
}
