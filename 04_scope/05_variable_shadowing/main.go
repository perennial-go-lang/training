package main

import "fmt"

// THIS IS A BAD PROGRAMMING PRACTICE. THIS EXAMPLE IS FOR ILLUSTRATION PURPOSES ONLY. NEVER DO THIS

func max(x int) int {
	return 30 + x
}

func main() {
	max := max(24)
	fmt.Println(max)
	//v := max(24) // max isn't a function anymore
}
