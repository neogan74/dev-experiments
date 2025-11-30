package _872_MaximumNumberofK_DivisibleComponents

import "testing"

func TestMaxKDivisibleComponents(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  [][]int
		values []int
		k      int
		want   int
	}{
		{
			name:   "example 1",
			n:      5,
			edges:  [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
			values: []int{1, 8, 1, 4, 4},
			k:      6,
			want:   2,
		},
		{
			name:   "example 2",
			n:      7,
			edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			values: []int{3, 0, 6, 1, 5, 2, 1},
			k:      3,
			want:   3,
		},
		{
			name:   "all divisible when k is one",
			n:      4,
			edges:  [][]int{{0, 1}, {1, 2}, {2, 3}},
			values: []int{5, 1, 2, 3},
			k:      1,
			want:   4,
		},
		{
			name:   "single node tree",
			n:      1,
			edges:  [][]int{},
			values: []int{6},
			k:      3,
			want:   1,
		},
		{
			name:   "no removable edge still valid",
			n:      2,
			edges:  [][]int{{0, 1}},
			values: []int{2, 3},
			k:      5,
			want:   1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxKDivisibleComponents(tt.n, tt.edges, tt.values, tt.k); got != tt.want {
				t.Fatalf("maxKDivisibleComponents() = %d, want %d", got, tt.want)
			}
		})
	}
}
