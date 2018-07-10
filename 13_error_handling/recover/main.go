package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	x := []int{}
	fmt.Println(x[0])
	// panic("PANIC")
}
