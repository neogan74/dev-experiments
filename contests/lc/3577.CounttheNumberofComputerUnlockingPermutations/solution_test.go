package main

import "testing"

func TestCountUnlockingPermutations(t *testing.T) {
	tests := []struct {
		name       string
		complexity []int
		want       int
	}{
		{
			name:       "example valid",
			complexity: []int{1, 2, 3},
			want:       2,
		},
		{
			name:       "example invalid",
			complexity: []int{3, 3, 3, 4, 4, 4},
			want:       0,
		},
		{
			name:       "just enough increase",
			complexity: []int{1, 2},
			want:       1, // only [0,1]
		},
		{
			name:       "fails when root not unique minimum",
			complexity: []int{2, 1, 3},
			want:       0,
		},
		{
			name:       "strictly increasing allows factorial",
			complexity: []int{1, 2, 5, 7, 8},
			want:       24, // (5-1)! = 24
		},
		{
			name:       "moderate size checks modulo logic",
			complexity: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			want:       3628800, // (11-1)! = 10!
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := countUnlockingPermutations(tc.complexity); got != tc.want {
				t.Fatalf("countUnlockingPermutations(%v) = %d, want %d", tc.complexity, got, tc.want)
			}
		})
	}
}
