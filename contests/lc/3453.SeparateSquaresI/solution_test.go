package separatesquaresi

import (
	"math"
	"testing"
)

func equal(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}

func Test_separateSquares(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		squares [][]int
		want    float64
	}{
		{
			name: "Test Case 1",
			squares: [][]int{
				{0, 0, 1},
				{2, 2, 1},
			},
			want: 1.0,
		},
		{
			name: "Test Case 2",
			squares: [][]int{
				{0, 0, 2},
				{1, 1, 1},
			},
			want: 1.16667,
		},
		// Additional test cases can be added here.
		{
			name: "Single square",
			squares: [][]int{
				{0, 0, 3},
			},
			want: 1.5,
		},
		{
			name: "Overlapping squares",
			squares: [][]int{
				{0, 0, 2},
				{0, 1, 2},
			},
			want: 1.5,
		},
		{
			name: "Edge case - All squares Same level",
			squares: [][]int{
				{0, 10, 2},
				{5, 10, 3},
				{10, 10, 4},
			},
			want: 11.61111111111111,
		},
		{
			name:    "Many Squares",
			squares: generateManySquares(1000),
			want:    499.99999999999994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := separateSquares(tt.squares)
			if !equal(got, tt.want) {
				t.Errorf("separateSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateManySquares(n int) [][]int {
	squares := make([][]int, n)
	for i := 0; i < n; i++ {
		squares[i] = []int{0, i, 1}
	}
	return squares
}
