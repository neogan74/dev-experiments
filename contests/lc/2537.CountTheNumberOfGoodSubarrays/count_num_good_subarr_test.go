package main

import "testing"

func TestCountGood(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{[]int{1, 1, 1, 1, 1}, 10, 1},
		{[]int{3, 1, 4, 3, 2, 2, 4}, 2, 4},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 1, 2, 2, 1, 1}, 3, 3},
	}

	for _, tt := range tests {
		result := countGood(tt.nums, tt.k)
		if result != tt.expected {
			t.Errorf("countGood(%v, %d) = %d; want %d", tt.nums, tt.k, result, tt.expected)
		}
	}
}
