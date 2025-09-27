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

func BenchmarkLargestTriangleArea(b *testing.B) {
	cases := []struct {
		name   string
		points [][]int
	}{
		{
			name:   "small",
			points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
		},
		{
			name:   "grid_5x5",
			points: benchmarkGridPoints(5),
		},
		{
			name:   "grid_7x7",
			points: benchmarkGridPoints(7),
		},
	}

	for _, tc := range cases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				largestTriangleArea(tc.points)
			}
		})
	}
}

func benchmarkGridPoints(n int) [][]int {
	points := make([][]int, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			points = append(points, []int{i, j})
		}
	}
	return points
}
