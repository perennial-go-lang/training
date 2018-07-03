package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Hello!",
		"Bye!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

}