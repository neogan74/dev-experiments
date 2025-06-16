package _3337_total_characters_in_string_after_transfor2

import "testing"

func TestLengthAfterTransformations(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        int
		nums     []int
		expected int
	}{
		{
			name:     "Base case",
			s:        "azbk",
			t:        1,
			nums:     makeUniformNums(2),
			expected: 8, // z→ab, a→bc, b→cd, k→lm
		},
		{
			name:     "No transformation",
			s:        "abc",
			t:        0,
			nums:     makeUniformNums(10),
			expected: 3,
		},
		{
			name:     "Only z case, t=1",
			s:        "z",
			t:        1,
			nums:     makeUniformNums(2),
			expected: 2, // z → a b
		},
		{
			name:     "Only z case, t=2",
			s:        "z",
			t:        2,
			nums:     makeUniformNums(2),
			expected: 4, // z → ab → bc + cd → total 4 chars
		},
		{
			name:     "Large t but simple case",
			s:        "a",
			t:        1_000_000_000,
			nums:     makeUniformNums(0), // no transformations
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthAfterTransformations(tc.s, tc.t, tc.nums)
			if got != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, got)
			}
		})
	}
}

func makeUniformNums(x int) []int {
	nums := make([]int, 26)
	for i := range nums {
		nums[i] = x
	}
	return nums
}
