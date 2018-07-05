package main

import (
	"fmt"
	mpkg "github.com/perennial-go-training/training/04_scope/04_visibility/mypackage"
)

func main() {
	fmt.Println(mpkg.MyName)
	mpkg.PrintCompanyName()
}
