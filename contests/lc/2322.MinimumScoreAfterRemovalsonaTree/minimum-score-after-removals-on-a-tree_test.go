package minimumscoreafterremovalsonatree

import (
	"testing"
)

func TestMinimumScore(t *testing.T) {
	tests := []struct {
		nums   []int
		edges  [][]int
		expect int
	}{
		{
			nums:   []int{1, 5, 5, 4, 11},
			edges:  [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}},
			expect: 9,
		},
		{
			nums:   []int{5, 5, 2, 4, 4, 2},
			edges:  [][]int{{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3}},
			expect: 0,
		},
	}

	for i, tt := range tests {
		got := minimumScore(tt.nums, tt.edges)
		if got != tt.expect {
			t.Errorf("Test %d: minimumScore(%v, %v) = %d; want %d", i+1, tt.nums, tt.edges, got, tt.expect)
		}
	}
}
