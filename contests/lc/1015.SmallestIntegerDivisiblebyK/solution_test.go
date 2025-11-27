package main

import "testing"

func TestSmallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		k        int
		expected int
	}{
		{k: 1, expected: 1},
		{k: 2, expected: -1},
		{k: 3, expected: 3},
		{k: 7, expected: 6},   // 111111
		{k: 17, expected: 16}, // known value
		{k: 19, expected: 18}, // known value
		{k: 9901, expected: 12},
	}

	for _, tt := range tests {
		if got := smallestRepunitDivByK(tt.k); got != tt.expected {
			t.Errorf("smallestRepunitDivByK(%d) = %d, want %d", tt.k, got, tt.expected)
		}
	}
}
