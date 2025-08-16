package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("when str is 'a' ", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertRepeat(t, got, want)
	})
}
func assertRepeat(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
