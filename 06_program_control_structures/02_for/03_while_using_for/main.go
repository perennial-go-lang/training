package main

import "fmt"

func main() {
	sum := 1

	//there is no while in GoLang, we need to use 'for' instead
	for sum < 1000 {
		sum += sum
	}

	//notice the difference between "for without init and post" vs "for as while"?
	//try implementing while(1) using for

	fmt.Println(sum)
}
