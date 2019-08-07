package main

import "fmt"

func main() {
	count := 0

	//for init; condition; post { }
	for i := 0; i < 10; i++ {
		count += i
	}
	fmt.Println(count)
}