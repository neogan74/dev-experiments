package _3355_zero_array_tranform

import "testing"

func Test_isZeroArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 0, 1},
			queries:  [][]int{{0, 2}},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 3, 2, 1},
			queries:  [][]int{{1, 3}, {0, 2}},
			expected: false,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			queries:  [][]int{},
			expected: true,
		},
		{
			name:     "Insufficient decrements",
			nums:     []int{2, 2, 2},
			queries:  [][]int{{0, 1}, {1, 2}},
			expected: false,
		},
		{
			name:     "Exact decrements",
			nums:     []int{2, 2, 2},
			queries:  [][]int{{0, 2}, {0, 2}},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isZeroArray(tc.nums, tc.queries)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

// TODO 1: add benchmark
