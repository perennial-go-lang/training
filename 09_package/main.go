package main

import (
	"fmt"
	"github.com/perennial-go-training/training/09_package/numutil"
	"github.com/perennial-go-training/training/09_package/stringutil"
)

func main() {
	message := "Hello, World!"
	reverse := stringutil.Reverse(message)
	fmt.Println(reverse)
	fmt.Println(numutil.OddOrEven(23736493284694))
	fmt.Println(numutil.OddOrEven(2123979120837921737))
}
