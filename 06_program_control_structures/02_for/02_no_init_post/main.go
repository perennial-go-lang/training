package main

import "fmt"

func main() {
	sum := 1

	//for init; condition; post { }
	//we can also use for without init and post, using just the condition

	for ; sum < 10; {
		sum += sum
	}
	fmt.Println(sum)
}
