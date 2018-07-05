package main

import "fmt"

func main() {
	var (
		number int
		fact uint64
	)
	fmt.Println("Enter a number")
	fmt.Scan(&number)

	f := factorial()

	for i := 1; i <= number; i++ {
		fact = f(uint64(i))
	}

	fmt.Print("Factorial of ", number, " is ", fact)
}

func factorial() func(n uint64)(uint64) {
	var a,b uint64 = 1, 1
	return func(n uint64)(uint64) {
		if n > 1 {
			a, b = b, n * b
		} else {
			return 1
		}
		return b
	}
}

