package solution

import "testing"

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example1",
			grid: [][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			want: 8,
		},
		{
			name: "example2",
			grid: [][]int{
				{3, 2},
				{1, 0},
			},
			want: 0,
		},
		{
			name: "allNegative",
			grid: [][]int{
				{-1, -1},
				{-2, -3},
			},
			want: 4,
		},
		{
			name: "singleRow",
			grid: [][]int{
				{5, 1, 0, -1, -2},
			},
			want: 2,
		},
		{
			name: "singleCol",
			grid: [][]int{
				{2},
				{1},
				{0},
				{-1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countNegatives(tt.grid)
			if got != tt.want {
				t.Fatalf("countNegatives() = %d, want %d", got, tt.want)
			}
		})
	}
}
