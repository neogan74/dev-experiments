package solution

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "example1",
			points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			want:   7,
		},
		{
			name:   "example2",
			points: [][]int{{3, 2}, {-2, 2}},
			want:   5,
		},
		{
			name:   "singlePoint",
			points: [][]int{{0, 0}},
			want:   0,
		},
		{
			name:   "mixedMoves",
			points: [][]int{{0, 0}, {1, 1}, {1, 2}},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minTimeToVisitAllPoints(tt.points)
			if got != tt.want {
				t.Fatalf("minTimeToVisitAllPoints() = %d, want %d", got, tt.want)
			}
		})
	}
}
