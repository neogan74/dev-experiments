package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Andrei", "English")
		want := "Hello, Andrei!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, GO!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Diego", "Spanish")
		want := "Hola, Diego!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Emma", "French")
		want := "Bonjour, Emma!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
