package main

import (
	"errors"
	"fmt"
	"log"

	. "github.com/rohitpavaskar/training/13_error_handling/main/exception"
)

func main() {
	a, b := 4, 0
	result, err := Divide(a, b)
	if err != nil {
		log.Fatalln("aaa")
	}

	fmt.Println(result)
}

func Divide(a, b int) (int, Exception) {
	if b == 0 {
		return 0, errors.New("ssssss")
	}
	return a / b, nil
}
