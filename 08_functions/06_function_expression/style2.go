package main

import "fmt"

func main() {
	greet := makeGreeter("Ajitem")
	greeting := greet()

	fmt.Println(greeting)
	fmt.Printf("%T", greet)

}

func makeGreeter(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}
