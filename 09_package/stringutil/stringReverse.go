package stringutil

import (
	"strings"
)

func StringReverse(str string) string {
	strRev := make([]string, 0)

	for i := len(str) - 1; i >= 0; i-- {
		strRev = append(strRev, string(str[i]))
	}
	return strings.Join(strRev, "")
}
