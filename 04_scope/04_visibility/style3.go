package main

import (
	"fmt"
	. "github.com/perennial-go-training/training/04_scope/04_visibility/mypackage"
)

// Be careful when importing a package this way to avoid variable shadowing

func main() {
	fmt.Println(MyName)
	PrintCompanyName()
}
