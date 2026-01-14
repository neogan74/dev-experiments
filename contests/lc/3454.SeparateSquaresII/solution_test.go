package main

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}

func TestSeparateSquares(t *testing.T) {
	tests := []struct {
		name    string
		squares [][]int
		want    float64
	}{
		{
			name:    "example1",
			squares: [][]int{{0, 0, 1}, {2, 2, 1}},
			want:    1.0,
		},
		{
			name:    "example2",
			squares: [][]int{{0, 0, 2}, {1, 1, 1}},
			want:    1.0,
		},
		{
			name:    "single_square",
			squares: [][]int{{5, 10, 4}},
			want:    12.0,
		},
		{
			name:    "identical_overlap",
			squares: [][]int{{0, 0, 2}, {0, 0, 2}},
			want:    1.0,
		},
		{
			name:    "gap_between",
			squares: [][]int{{0, 0, 1}, {0, 2, 1}},
			want:    1.0,
		},
		{
			name:    "rect_union",
			squares: [][]int{{0, 0, 2}, {1, 0, 2}},
			want:    1.0,
		},
		{
			name:    "different_heights",
			squares: [][]int{{0, 0, 4}, {0, 4, 2}},
			want:    2.5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := separateSquares(tc.squares)
			if !almostEqual(got, tc.want, 1e-6) {
				t.Fatalf("separateSquares() = %.10f, want %.10f", got, tc.want)
			}
		})
	}
}
