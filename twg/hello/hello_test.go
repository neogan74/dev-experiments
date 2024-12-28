package hello

import "testing"

func TestHello(t *testing.T) {
	name := "Go"
	got := Hello(name)
	want := "Hello, Go!"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

}
