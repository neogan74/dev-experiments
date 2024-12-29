package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAria(t *testing.T) {
	checkResult := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.g want %.g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0
		checkResult(t, got, want)
	})
	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		checkResult(t, got, want)
	})
}
