package numutil

func OddOrEven(num int) (result string) {
	if (num % 2) == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	return
}
