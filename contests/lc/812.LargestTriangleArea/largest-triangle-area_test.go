package largesttrianglearea

import (
	"math"
	"testing"
)

func TestLargestTriangleArea(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		points [][]int
		want   float64
	}{
		{
			name:   "example from problem",
			points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			want:   2.0,
		},
		{
			name:   "colinear points",
			points: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
			want:   0.0,
		},
		{
			name:   "negative coordinates",
			points: [][]int{{-1, 0}, {0, 0}, {0, 3}, {2, 0}},
			want:   4.5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := largestTriangleArea(tt.points)
			if diff := math.Abs(got - tt.want); diff > 1e-9 {
				t.Fatalf("largestTriangleArea() = %v, want %v (diff %.12f)", got, tt.want, diff)
			}
		})
	}
}
