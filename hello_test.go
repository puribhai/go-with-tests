package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("prints 'Hello World' on passing default empty string", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"
		assertCorrectMessage(t, got, want)

	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
