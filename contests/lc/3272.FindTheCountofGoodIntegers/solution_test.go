package main

import "testing"

func TestCountGoodIntegers(t *testing.T) {
	tests := []struct {
		n, k     int
		expected int64
	}{
		{1, 1, 9},
		{2, 1, 90},
		{2, 11, 9},
		{3, 1, 252},
		{4, 2, 90},
	}

	for _, tt := range tests {
		result := countGoodIntegers(tt.n, tt.k)
		if result != tt.expected {
			t.Errorf("countGoodIntegers(%d, %d) = %d; want %d", tt.n, tt.k, result, tt.expected)
		}
	}
}
