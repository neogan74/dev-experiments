package main

import "testing"

func TestCountPartitionsExamples(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{9, 4, 1, 3, 7}, k: 4, want: 6},
		{nums: []int{3, 3, 4}, k: 0, want: 2},
	}

	for _, tc := range cases {
		got := countPartitions(tc.nums, tc.k)
		if got != tc.want {
			t.Fatalf("nums=%v k=%d: want %d, got %d", tc.nums, tc.k, tc.want, got)
		}
	}
}

func TestCountPartitionsAdditional(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		// k large: all partitions are valid -> 2^(n-1) for n=3
		{nums: []int{1, 5, 2}, k: 10, want: 4},
		// Tight window: only segments with max-min <= 1 allowed
		{nums: []int{1, 3, 2}, k: 1, want: 2},
		// All equal with k=0: every split valid -> 2^(n-1) for n=3
		{nums: []int{5, 5, 5}, k: 0, want: 4},
	}

	for _, tt := range tests {
		got := countPartitions(tt.nums, tt.k)
		if got != tt.want {
			t.Fatalf("nums=%v k=%d: want %d, got %d", tt.nums, tt.k, tt.want, got)
		}
	}
}
