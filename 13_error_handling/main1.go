package main

import (
	"fmt"
	"errors"
	"log"
)

type Exception interface {
	//Action()
	GetError() error
}

type DivisionError struct {
	Error error
	Action func()
}

//func (e DivisionError) Action() {
//	fmt.Println("Some action")
//}

func (e DivisionError) GetError() error {
	defer e.Action()
	return e.Error
}

var DivideByZero = DivisionError{errors.New("Cannot Divide by Zero"), func() {
}}

func main() {
	var num1, num2, quotient int
	var err Exception

	num1 = 4
	num2 = 0

	quotient, err = Divide(num1, num2)

	if err != nil {
		log.Fatalln(err.GetError())
	}

	fmt.Println(quotient)
}

func Divide(a, b int) (int, Exception) {
	if b == 0 {
		return 0, DivideByZero
	}
	return  a / b, nil
}
