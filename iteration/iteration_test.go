package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("when str is 'a' ", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertRepeat(t, got, want)
	})

	t.Run("when str is 'ab'", func(t *testing.T) {
		got := Repeat("ab")
		want := "ababababab"
		assertRepeat(t, got, want)
	})

	t.Run("when str is empty", func(t *testing.T) {
		got := Repeat("")
		want := ""
		assertRepeat(t, got, want)
	})
}
func assertRepeat(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
