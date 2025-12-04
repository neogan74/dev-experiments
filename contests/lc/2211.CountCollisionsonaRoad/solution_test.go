package main

import "testing"

func TestCountCollisions(t *testing.T) {
	tests := []struct {
		directions string
		want       int
	}{
		{"RLRSLL", 5},
		{"LLRR", 0},
		{"S", 0},
		{"L", 0},
		{"R", 0},
		{"SRRS", 2},
		{"RRSLL", 4},
	}

	for _, tt := range tests {
		if got := countCollisions(tt.directions); got != tt.want {
			t.Fatalf("countCollisions(%q) = %d, want %d", tt.directions, got, tt.want)
		}
	}
}

func BenchmarkCountCollisions(b *testing.B) {
	input := "RRSLLSRRSLRSLLRSRSLLRRSRLSLRRSLLSRS"
	for i := 0; i < b.N; i++ {
		_ = countCollisions(input)
	}
}
