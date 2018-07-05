package numutil

func OddOrEven(x int) string {
	return []string{"Even", "Odd"}[x & 1]
}