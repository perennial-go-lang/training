package main

import "fmt"

func main() {
	language := "Go"

	switch {
	case len(language) == 2:
		fmt.Println("What language has a two character name?")
	case language == "Go":
		fmt.Println("I am learning Go")
	case language == "C":
		fmt.Println("I am learning C")
	case language == "C++":
		fmt.Println("I am learning C++")
	case language == "PHP":
		fmt.Println("I am learning PHP")
	case language == "Java":
		fmt.Println("I am learning Java")
	}
}
