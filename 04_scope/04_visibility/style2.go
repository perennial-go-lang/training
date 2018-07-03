package main

import (
	"fmt"
	mpkg "github.com/perennial-go-training/training/05_scopey/mypackage"
)

func main() {
	fmt.Println(mpkg.MyName)
	mpkg.PrintCompanyName()
}
