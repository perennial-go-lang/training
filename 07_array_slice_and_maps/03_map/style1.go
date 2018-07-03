package main

import "fmt"

func main() {
	var greeting map[string]string

	fmt.Println(greeting)
	fmt.Println(greeting == nil)

	greeting["person1"] = "Hi!"
	greeting["person2"] = "Hello!"

	fmt.Println(greeting)
}
