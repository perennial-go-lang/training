package numutil

func OddOrEven(num int) (res string) {
	switch {
	case num%2 == 0:
		return "Even"
	default:
		return "Odd"
	}
}
