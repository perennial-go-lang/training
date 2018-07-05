package stringutil

func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(str) - 1; i < len(str)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
