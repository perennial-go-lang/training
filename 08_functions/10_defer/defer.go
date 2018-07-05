package main

import "fmt"

func hello()  {
	fmt.Print("Hello")
}

func world() {
	fmt.Print("World")
}

func main() {
	defer hello()
	func() {
		fmt.Print("Yahoo")
	}()
	defer world()
}