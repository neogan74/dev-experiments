package main

import "testing"

func TestMaximizeSquareHoleAreaExamples(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		h    []int
		v    []int
		want int
	}{
		{
			name: "example1",
			n:    2,
			m:    1,
			h:    []int{2, 3},
			v:    []int{2},
			want: 4,
		},
		{
			name: "example2",
			n:    1,
			m:    1,
			h:    []int{2},
			v:    []int{2},
			want: 4,
		},
		{
			name: "example3",
			n:    2,
			m:    3,
			h:    []int{2, 3},
			v:    []int{2, 4},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximizeSquareHoleArea(tt.n, tt.m, append([]int(nil), tt.h...), append([]int(nil), tt.v...))
			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaximizeSquareHoleAreaExtra(t *testing.T) {
	if got := maximizeSquareHoleArea(5, 5, []int{2, 4, 6}, []int{3, 5}); got != 4 {
		t.Fatalf("got %d, want %d", got, 4)
	}
	if got := maximizeSquareHoleArea(5, 5, []int{2, 3, 4}, []int{2, 4, 5}); got != 9 {
		t.Fatalf("got %d, want %d", got, 9)
	}
}
