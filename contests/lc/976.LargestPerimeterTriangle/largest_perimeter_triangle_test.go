package main

import "testing"

func TestLargestPerimeter(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		nums     []int
		expected int
	}{
		"valid_triangle": {
			nums:     []int{2, 1, 2},
			expected: 5,
		},
		"no_triangle": {
			nums:     []int{1, 2, 1, 10},
			expected: 0,
		},
		"skip_non_qualifying_largest": {
			nums:     []int{3, 6, 2, 3},
			expected: 8,
		},
		"all_sides_equal": {
			nums:     []int{6, 6, 6},
			expected: 18,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := largestPerimeter(append([]int(nil), tc.nums...)); got != tc.expected {
				t.Fatalf("largestPerimeter(%v) = %d, want %d", tc.nums, got, tc.expected)
			}
		})
	}
}
