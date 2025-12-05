package main

import "testing"

func TestCountPartitions(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 10, 3, 7, 6}, 4},
		{[]int{1, 2, 2}, 0},
		{[]int{2, 4, 6, 8}, 3},
	}

	for _, tt := range tests {
		if got := countPartitions(tt.nums); got != tt.expected {
			t.Fatalf("countPartitions(%v) = %d, want %d", tt.nums, got, tt.expected)
		}
	}
}
