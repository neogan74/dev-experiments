package main

import "testing"

func TestMaxDotProduct(t *testing.T) {
	cases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "example1",
			nums1: []int{2, 1, -2, 5},
			nums2: []int{3, 0, -6},
			want:  18,
		},
		{
			name:  "example2",
			nums1: []int{3, -2},
			nums2: []int{2, -6, 7},
			want:  21,
		},
		{
			name:  "example3",
			nums1: []int{-1, -1},
			nums2: []int{1, 1},
			want:  -1,
		},
		{
			name:  "single_negative_best",
			nums1: []int{-5, -1},
			nums2: []int{2, 3},
			want:  -2,
		},
		{
			name:  "single_pair",
			nums1: []int{7},
			nums2: []int{-3},
			want:  -21,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxDotProduct(tc.nums1, tc.nums2)
			if got != tc.want {
				t.Fatalf("maxDotProduct(%v, %v) = %d, want %d", tc.nums1, tc.nums2, got, tc.want)
			}
		})
	}
}
