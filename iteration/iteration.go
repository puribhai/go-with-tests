package iteration

// Repeat returns a new string consisting of the input character repeated 5 times
func Repeat(str string) string {
	var ans string
	for i := 0; i < 5; i++ {
		ans += str
	}
	return ans
}
