package main

import (
	"fmt"

	evenOddPack "github.com/ankitkumbhar/training/09_package/numutil"
	strRevPack "github.com/ankitkumbhar/training/09_package/stringutil"
)

func main() {
	strReverse := strRevPack.StringReverse("abcdefghijklmn")
	fmt.Println(strReverse)
	result := evenOddPack.EvenOdd(11)
	fmt.Printf("Number is " + result)
}
