package maximumuniquesubarraysumafterdeletion

import (
	"testing"
)

// TestMaximumUniqueSubarraySumAfterDeletion tests the maximumUniqueSubarraySumAfterDeletion function
// with various input cases to ensure it correctly computes the maximum sum of a unique subarray
// after deleting duplicates. The test covers edge cases such as empty arrays, arrays with all
// unique elements, arrays with all duplicates, and mixed cases to validate the correctness of the
// implementation.
func TestMaximumUniqueSubarraySumAfterDeletion(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{4, 2, 4, 5, 6}, 17},
		{[]int{1, 2, 1, 3, 4, 2, 3}, 10},
		{[]int{5, 5, 5, 5, 5}, 5},
		{[]int{1, 2, 3, 2, 1, 4, 5}, 15},
		{[]int{2, 2, 2, 2, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 1, 2, 3, 4, 5}, 15},
	}

	for _, tt := range tests {
		result := maximumUniqueSubarraySumAfterDeletion(tt.nums)
		if result != tt.expected {
			t.Errorf("maximumUniqueSubarraySumAfterDeletion(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
