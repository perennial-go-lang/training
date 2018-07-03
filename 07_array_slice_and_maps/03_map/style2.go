package main

import "fmt"

func main() {

	var greeting = make(map[string]string)

	greeting["person1"] = "Hi!"
	greeting["person2"] = "Hello!"

	fmt.Println(greeting)
}