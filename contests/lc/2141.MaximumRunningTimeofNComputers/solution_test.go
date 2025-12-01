package maximumrunningtimeofncomputers

import "testing"

func TestMaxRunTime(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		batteries []int
		want      int64
	}{
		{
			name:      "two computers equal batteries",
			n:         2,
			batteries: []int{3, 3, 3},
			want:      4,
		},
		{
			name:      "two computers small batteries",
			n:         2,
			batteries: []int{1, 1, 1, 1},
			want:      2,
		},
		{
			name:      "three computers mixed batteries",
			n:         3,
			batteries: []int{10, 10, 3, 5},
			want:      8,
		},
		{
			name:      "all batteries sufficient",
			n:         2,
			batteries: []int{5, 4, 3},
			want:      6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRunTime(tt.n, tt.batteries); got != tt.want {
				t.Fatalf("maxRunTime(%d, %v) = %d, want %d", tt.n, tt.batteries, got, tt.want)
			}
		})
	}
}

func TestMaxRunTime2(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		batteries []int
		want      int64
	}{
		{
			name:      "two computers equal batteries",
			n:         2,
			batteries: []int{3, 3, 3},
			want:      4,
		},
		{
			name:      "two computers small batteries",
			n:         2,
			batteries: []int{1, 1, 1, 1},
			want:      2,
		},
		{
			name:      "three computers mixed batteries",
			n:         3,
			batteries: []int{10, 10, 3, 5},
			want:      8,
		},
		{
			name:      "all batteries sufficient",
			n:         2,
			batteries: []int{5, 4, 3},
			want:      6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRunTime2(tt.n, tt.batteries); got != tt.want {
				t.Fatalf("maxRunTime(%d, %v) = %d, want %d", tt.n, tt.batteries, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxRunTime(b *testing.B) {
	cases := []struct {
		name      string
		n         int
		batteries []int
	}{
		{
			name:      "small",
			n:         2,
			batteries: []int{3, 3, 3},
		},
		{
			name:      "medium",
			n:         3,
			batteries: []int{10, 10, 3, 5, 7, 9, 2},
		},
		{
			name:      "large",
			n:         5,
			batteries: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		},
	}

	for _, tt := range cases {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = maxRunTime(tt.n, tt.batteries)
			}
		})
	}
}

func BenchmarkMaxRunTime2(b *testing.B) {
	cases := []struct {
		name      string
		n         int
		batteries []int
	}{
		{
			name:      "small",
			n:         2,
			batteries: []int{3, 3, 3},
		},
		{
			name:      "medium",
			n:         3,
			batteries: []int{10, 10, 3, 5, 7, 9, 2},
		},
		{
			name:      "large",
			n:         5,
			batteries: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		},
	}

	for _, tt := range cases {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = maxRunTime2(tt.n, tt.batteries)
			}
		})
	}
}
