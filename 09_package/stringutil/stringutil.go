package stringutil

func ReverseString(source string) string {
	strLength := len(source)
	result := []rune(source)
	for i, j := 0, strLength-1; i < strLength/2; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
