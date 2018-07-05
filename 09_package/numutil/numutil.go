package numutil

func OddOrEven(number int) (result string) {
	if (number % 2) == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	return
}
