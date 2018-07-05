package numutil

func EvenOdd(num int) (result string) {
	if (num % 2) == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	return result
}
