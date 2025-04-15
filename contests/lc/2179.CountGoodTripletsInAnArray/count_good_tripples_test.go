package main

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected int64
	}{
		{
			nums1:    []int{2, 0, 1, 3},
			nums2:    []int{0, 1, 2, 3},
			expected: 1,
		},
		{
			nums1:    []int{4, 0, 1, 3, 2},
			nums2:    []int{4, 1, 0, 2, 3},
			expected: 4,
		},
	}

	for _, tt := range tests {
		result := countGoodTriplets(tt.nums1, tt.nums2)
		if result != tt.expected {
			t.Errorf("countGoodTriplets(%v, %v) = %d; want %d", tt.nums1, tt.nums2, result, tt.expected)
		}
	}
}
