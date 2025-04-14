package main

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	tests := []struct {
		arr      []int
		a, b, c  int
		expected int
	}{
		{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
		{[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
		{[]int{1, 2, 3}, 10, 10, 10, 1},
		{[]int{1, 2, 3, 4, 5}, 1, 1, 1, 0},
	}

	for _, tt := range tests {
		result := countGoodTriplets(tt.arr, tt.a, tt.b, tt.c)
		if result != tt.expected {
			t.Errorf("countGoodTriplets(%v, %d, %d, %d) = %d; want %d",
				tt.arr, tt.a, tt.b, tt.c, result, tt.expected)
		}
	}
}

func TestCountGoodTripletsOptimized(t *testing.T) {
	tests := []struct {
		arr      []int
		a, b, c  int
		expected int
	}{
		{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
		{[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
		{[]int{1, 2, 3}, 10, 10, 10, 1},
		{[]int{1, 2, 3, 4, 5}, 1, 1, 1, 0},
	}

	for _, tt := range tests {
		result := countGoodTripletsOptimized(tt.arr, tt.a, tt.b, tt.c)
		if result != tt.expected {
			t.Errorf("countGoodTripletsOptimized(%v, %d, %d, %d) = %d; want %d",
				tt.arr, tt.a, tt.b, tt.c, result, tt.expected)
		}
	}
}
