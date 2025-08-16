package iteration

import "strings"

// Repeat returns a new string consisting of the input character repeated 5 times
func Repeat(str string) string {
	var ans strings.Builder
	for i := 0; i < 5; i++ {
		ans.WriteString(str)
	}
	return ans.String()
}
