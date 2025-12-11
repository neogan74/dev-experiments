package countcoveredbuildings

import "testing"

func TestCountCoveredBuildingsExamples(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		buildings [][]int
		want      int
	}{
		{
			name:      "example 1",
			n:         3,
			buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}},
			want:      1,
		},
		{
			name:      "example 2",
			n:         3,
			buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			want:      0,
		},
		{
			name:      "example 3",
			n:         5,
			buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}},
			want:      1,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := countCoveredBuildings(tc.n, tc.buildings); got != tc.want {
				t.Fatalf("countCoveredBuildings() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestCountCoveredBuildingsAdditional(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		buildings [][]int
		want      int
	}{
		{
			name:      "full 4x4 grid, four covered interiors",
			n:         4,
			buildings: fullGrid(4),
			want:      4,
		},
		{
			name:      "sparse corners only, none covered",
			n:         4,
			buildings: [][]int{{2, 2}, {2, 4}, {4, 2}, {4, 4}},
			want:      0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := countCoveredBuildings2(tc.n, tc.buildings); got != tc.want {
				t.Fatalf("countCoveredBuildings() = %d, want %d", got, tc.want)
			}
		})
	}
}

func fullGrid(n int) [][]int {
	out := make([][]int, 0, n*n)
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			out = append(out, []int{r, c})
		}
	}
	return out
}
