package lengthoflongestvshapeddiagonalsegment

import "testing"

func Test_lenOfVDiagonal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid [][]int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{
				{2, 2, 1, 2, 2},
				{2, 0, 2, 2, 0},
				{2, 0, 1, 1, 0},
				{1, 0, 2, 2, 2},
				{2, 0, 0, 2, 2},
			},
			want: 5,
		},
		{
			name: "example 2",
			grid: [][]int{
				{2, 2, 2, 2, 2},
				{2, 0, 2, 2, 0},
				{2, 0, 1, 1, 0},
				{1, 0, 2, 2, 2},
				{2, 0, 0, 2, 2},
			},
			want: 4,
		},
		{
			name: "example 3",
			grid: [][]int{
				{1, 2, 2, 2, 2},
				{2, 2, 2, 2, 0},
				{2, 0, 0, 0, 0},
				{0, 0, 2, 2, 2},
				{2, 0, 0, 2, 0},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lenOfVDiagonal(tt.grid)
			if got != tt.want {
				t.Errorf("lenOfVDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
