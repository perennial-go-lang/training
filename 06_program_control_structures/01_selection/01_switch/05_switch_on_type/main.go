package main

import "fmt"

func switchOnType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Variable is integer")
	case string:
		fmt.Println("Variable is string")
	case bool:
		fmt.Println("Variable is boolean")
	case float64:
		fmt.Println("Variable is a floating point")
	case rune:
		fmt.Println("Variable is a character (rune)")
	default:
		fmt.Println("Variable is of unknown type")
	}
}

func main() {
	//variable := "ABC"
	//variable := 123
	//variable := true
	//variable := 1.23
	variable := 'a'

	switchOnType(variable)
}
