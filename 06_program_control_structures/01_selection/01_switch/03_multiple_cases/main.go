package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	switch name {
	case "Ajitem", "Babulal", "Rohit", "Ankit":
		fmt.Println("Hello, Guys!")
	case "Soni", "Kalpashree":
		fmt.Println("Hello, Girls!")
	default:
		fmt.Println("Hello, " + name + " you are not a part of GoLang Training. Would you like to join us?")
	}
}