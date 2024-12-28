package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Andrei")
		want := "Hello, Andrei!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, GO!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
