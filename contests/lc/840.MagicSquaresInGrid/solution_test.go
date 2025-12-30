package magicsquaresingrid

import (
	"testing"
)

func TestNumMagicSquaresInside(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{4, 3, 8, 4},
				{9, 5, 1, 9},
				{2, 7, 6, 2},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{8},
			},
			expected: 0,
		},
		{
			name: "Small Grid",
			grid: [][]int{
				{4, 3},
				{9, 5},
			},
			expected: 0,
		},
		{
			name: "Valid Magic Square",
			grid: [][]int{
				{2, 7, 6},
				{9, 5, 1},
				{4, 3, 8},
			},
			expected: 1,
		},
		{
			name: "Invalid Magic Square (duplicate numbers)",
			grid: [][]int{
				{2, 7, 6},
				{9, 5, 1},
				{4, 3, 9},
			},
			expected: 0,
		},
		{
			name: "Grid with Multiple Magic Squares",
			grid: [][]int{
				{4, 3, 8, 4, 2},
				{9, 5, 1, 9, 7},
				{2, 7, 6, 2, 8},
				{4, 3, 8, 4, 2},
				{9, 5, 1, 9, 7},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numMagicSquaresInside(tt.grid)
			if result != tt.expected {
				t.Errorf("numMagicSquaresInside() = %v, want %v", result, tt.expected)
			}
		})
	}
}
