package main

import "fmt"

func main() {

	greeting := make(map[string]string)

	greeting["person1"] = "Hi!"
	greeting["person2"] = "Hello!"

	fmt.Println(greeting)
}