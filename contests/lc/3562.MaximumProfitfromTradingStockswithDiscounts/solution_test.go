package maximumprofitfromtradingstockswithdiscounts

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		present   []int
		future    []int
		hierarchy [][]int
		budget    int
		expected  int
	}{
		{
			name:      "Example 1",
			n:         2,
			present:   []int{1, 2},
			future:    []int{4, 3},
			hierarchy: [][]int{{1, 2}},
			budget:    3,
			expected:  5,
		},
		{
			name:      "Example 2",
			n:         3,
			present:   []int{5, 2, 3},
			future:    []int{8, 5, 6},
			hierarchy: [][]int{{1, 2}, {2, 3}},
			budget:    7,
			expected:  12,
		},
		{
			name:      "No hierarchy",
			n:         1,
			present:   []int{5},
			future:    []int{10},
			hierarchy: [][]int{},
			budget:    5,
			expected:  5,
		},
		{
			name:      "No profit possible",
			n:         1,
			present:   []int{10},
			future:    []int{5},
			hierarchy: [][]int{},
			budget:    10,
			expected:  0,
		},
		{
			name:      "Multiple children",
			n:         4,
			present:   []int{10, 5, 3, 2},
			future:    []int{20, 10, 8, 6},
			hierarchy: [][]int{{1, 2}, {1, 3}, {1, 4}},
			budget:    10,
			expected:  14,
		},
		{
			name:      "Deep hierarchy",
			n:         5,
			present:   []int{10, 4, 4, 2, 2},
			future:    []int{20, 10, 10, 5, 5},
			hierarchy: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			budget:    10,
			expected:  22,
		},
		{
			name:      "Multiple roots",
			n:         3,
			present:   []int{10, 5, 5},
			future:    []int{20, 10, 10},
			hierarchy: [][]int{{1, 2}, {1, 3}},
			budget:    15,
			expected:  26,
		},
		{
			name:      "Budget too low",
			n:         2,
			present:   []int{10, 5},
			future:    []int{20, 10},
			hierarchy: [][]int{{1, 2}},
			budget:    5,
			expected:  5,
		},
		{
			name:      "Full discount chain",
			n:         4,
			present:   []int{8, 4, 2, 1},
			future:    []int{16, 8, 4, 2},
			hierarchy: [][]int{{1, 2}, {2, 3}, {3, 4}},
			budget:    15,
			expected:  19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.n, tt.present, tt.future, tt.hierarchy, tt.budget)
			if result != tt.expected {
				t.Errorf("maxProfit(%v, %v, %v, %v, %v) = %v; want %v",
					tt.n, tt.present, tt.future, tt.hierarchy, tt.budget, result, tt.expected)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.n, tt.present, tt.future, tt.hierarchy, tt.budget)
			if result != tt.expected {
				t.Errorf("maxProfit2(%v, %v, %v, %v, %v) = %v; want %v",
					tt.n, tt.present, tt.future, tt.hierarchy, tt.budget, result, tt.expected)
			}
		})
	}
}
