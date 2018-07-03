package main

import (
	"fmt"
	"github.com/perennial-go-training/training/05_scopey/mypackage"
)

func main() {
	fmt.Println(mypackage.MyName)
	mypackage.PrintCompanyName()
}
