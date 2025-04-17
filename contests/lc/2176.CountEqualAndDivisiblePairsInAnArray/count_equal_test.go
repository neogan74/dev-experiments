package _176_CountEqualAndDivisiblePairsInAnArray

import "testing"

func TestCountPairs(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 1, 2, 2, 2, 1, 3}, 2, 4},
		{[]int{1, 2, 3, 4}, 1, 0},
		{[]int{1, 1, 1, 1}, 1, 6}, // все пары
		{[]int{1, 2, 1, 2, 1}, 2, 3},
	}

	for _, tt := range tests {
		result := countPairs(tt.nums, tt.k)
		if result != tt.expected {
			t.Errorf("countPairs(%v, %d) = %d; want %d", tt.nums, tt.k, result, tt.expected)
		}
	}
}
