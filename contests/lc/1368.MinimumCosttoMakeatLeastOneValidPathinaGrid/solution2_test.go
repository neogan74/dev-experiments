package minimumcosttomakeatleastonevalidpathinagrid

import "testing"

func TestMinCost2(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{
				{1, 1, 1, 1},
				{2, 2, 2, 2},
				{1, 1, 1, 1},
				{2, 2, 2, 2},
			},
			want: 3,
		},
		{
			name: "example 2",
			grid: [][]int{
				{1, 1, 3},
				{3, 2, 2},
				{1, 1, 4},
			},
			want: 0,
		},
		{
			name: "example 3",
			grid: [][]int{
				{1, 2},
				{4, 3},
			},
			want: 1,
		},
		{
			name: "single cell",
			grid: [][]int{
				{4},
			},
			want: 0,
		},
		{
			name: "needs two changes",
			grid: [][]int{
				{2, 2},
				{2, 2},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost2(tt.grid); got != tt.want {
				t.Errorf("minCost2() = %d, want %d", got, tt.want)
			}
		})
	}
}
