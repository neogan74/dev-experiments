package _3362_zero_array_transformation_iii

import "testing"

func TestMaxRemoval(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 0, 1},
			queries:  [][]int{{0, 2}},
			expected: 0,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 3, 2, 1},
			queries:  [][]int{{1, 3}, {0, 2}},
			expected: -1,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			queries:  [][]int{},
			expected: 0,
		},
		{
			name:     "Exact decrements",
			nums:     []int{2, 2, 2},
			queries:  [][]int{{0, 2}, {0, 2}},
			expected: 0,
		},
		{
			name:     "Insufficient coverage",
			nums:     []int{2, 2, 2},
			queries:  [][]int{{0, 1}, {1, 2}},
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxRemoval(tc.nums, tc.queries)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
