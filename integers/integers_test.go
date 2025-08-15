package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("if a = 3, b = 4", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assertSum(t, got, want)
	})
	t.Run("if a = 0, b = 0", func(t *testing.T) {
		got := Add(0, 0)
		want := 0

		assertSum(t, got, want)
	})
	t.Run("if a = 30, b = -15", func(t *testing.T) {
		got := Add(30, -15)
		want := 15

		assertSum(t, got, want)
	})
}

func assertSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' but want '%d'", got, want)
	}
}
