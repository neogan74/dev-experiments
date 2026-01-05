package main

import "testing"

func TestMaxMatrixSum(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int64
	}{
		{
			name:   "example1",
			matrix: [][]int{{1, -1}, {-1, 1}},
			want:   4,
		},
		{
			name:   "example2",
			matrix: [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}},
			want:   16,
		},
		{
			name:   "odd_negatives_no_zero",
			matrix: [][]int{{-1, 2}, {3, 4}},
			want:   8,
		},
		{
			name:   "odd_negatives_with_zero",
			matrix: [][]int{{-1, 0}, {2, 3}},
			want:   6,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := maxMatrixSum(tc.matrix)
			if got != tc.want {
				t.Fatalf("maxMatrixSum() = %d, want %d", got, tc.want)
			}
		})
	}
}
