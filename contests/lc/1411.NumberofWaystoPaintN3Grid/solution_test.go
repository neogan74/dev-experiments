package numberofwaystopaintn3grid

import "testing"

func TestNumOfWays(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 12},
		{n: 2, want: 54},
		{n: 3, want: 246},
		{n: 5000, want: 30228214},
	}

	for _, tc := range tests {
		if got := numOfWays(tc.n); got != tc.want {
			t.Fatalf("numOfWays(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}
