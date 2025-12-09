package main

import "testing"

func TestCountSpecialTripletsExamples(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{6, 3, 6}, 1},
		{[]int{0, 1, 0, 0}, 1},
		{[]int{8, 4, 2, 8, 4}, 2},
	}

	for _, tc := range cases {
		if got := countSpecialTriplets(tc.nums); got != tc.want {
			t.Fatalf("nums=%v: got %d, want %d", tc.nums, got, tc.want)
		}
	}
}

func TestCountSpecialTripletsNoMatches(t *testing.T) {
	if got := countSpecialTriplets([]int{1, 2, 3, 4}); got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestCountSpecialTripletsMultiplePairs(t *testing.T) {
	nums := []int{6, 6, 3, 6, 6} // middle 3 can pair with two left and two right doubles
	if got := countSpecialTriplets(nums); got != 4 {
		t.Fatalf("nums=%v: got %d, want 4", nums, got)
	}
}

func TestSpecialTripletsWrapper(t *testing.T) {
	nums := []int{8, 4, 2, 8, 4}
	if got := specialTriplets(nums); got != 2 {
		t.Fatalf("specialTriplets: nums=%v: got %d, want 2", nums, got)
	}
}
