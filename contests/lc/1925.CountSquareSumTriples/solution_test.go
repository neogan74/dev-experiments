package countsquaresumtriples

import "testing"

func TestCountTriples(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"example_5", 5, 2},
		{"example_10", 10, 4},
		{"no_triples", 1, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriples(tt.n); got != tt.want {
				t.Fatalf("countTriples(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestCountTriples2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"example_5", 5, 2},
		{"example_10", 10, 4},
		{"no_triples", 1, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriples2(tt.n); got != tt.want {
				t.Fatalf("countTriples2(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestImplementationsMatch(t *testing.T) {
	for n := 1; n <= 30; n++ {
		a := countTriples(n)
		b := countTriples2(n)
		if a != b {
			t.Fatalf("countTriples mismatch at n=%d: %d vs %d", n, a, b)
		}
	}
}

func BenchmarkCountTriples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countTriples(250)
	}
}

func BenchmarkCountTriples2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countTriples2(250)
	}
}
