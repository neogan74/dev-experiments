package main

import "testing"

type testCase struct {
	bottomLeft [][]int
	topRight   [][]int
	expected   int64
}

func TestLargestSquareArea(t *testing.T) {
	tests := []testCase{
		{
			bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}},
			topRight:   [][]int{{3, 3}, {4, 4}, {6, 6}},
			expected:   1,
		},
		{
			bottomLeft: [][]int{{1, 1}, {1, 3}, {1, 5}},
			topRight:   [][]int{{5, 5}, {5, 7}, {5, 9}},
			expected:   4,
		},
		{
			bottomLeft: [][]int{{1, 1}, {2, 2}, {1, 2}},
			topRight:   [][]int{{3, 3}, {4, 4}, {3, 4}},
			expected:   1,
		},
		{
			bottomLeft: [][]int{{1, 1}, {3, 3}, {3, 1}},
			topRight:   [][]int{{2, 2}, {4, 4}, {4, 2}},
			expected:   0,
		},
		{
			bottomLeft: [][]int{{1, 1}, {1, 1}},
			topRight:   [][]int{{6, 4}, {4, 8}},
			expected:   9,
		},
	}

	for i, tc := range tests {
		got := largestSquareArea(tc.bottomLeft, tc.topRight)
		if got != tc.expected {
			t.Fatalf("case %d: expected %d, got %d", i, tc.expected, got)
		}
	}
}
