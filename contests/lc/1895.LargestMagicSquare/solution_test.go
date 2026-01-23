package main

import "testing"

func TestLargestMagicSquare(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example1",
			grid: [][]int{
				{7, 1, 4, 5, 6},
				{2, 5, 1, 6, 4},
				{1, 5, 4, 3, 2},
				{1, 2, 7, 3, 4},
			},
			want: 3,
		},
		{
			name: "example2",
			grid: [][]int{
				{5, 1, 3, 1},
				{9, 3, 3, 1},
				{1, 3, 3, 8},
			},
			want: 2,
		},
		{
			name: "single",
			grid: [][]int{{42}},
			want: 1,
		},
		{
			name: "no_magic_larger",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: 1,
		},
		{
			name: "rectangular_3x4",
			grid: [][]int{
				{2, 7, 6, 2},
				{9, 5, 1, 9},
				{4, 3, 8, 4},
			},
			want: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := largestMagicSquare(tc.grid); got != tc.want {
				t.Fatalf("largestMagicSquare() = %d, want %d", got, tc.want)
			}
		})
	}
}
