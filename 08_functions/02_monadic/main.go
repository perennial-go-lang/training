package main

import "fmt"

func main() {
	greet("Ajitem")
}

func greet(name string) {
	fmt.Println("Hello", name, "!")
}