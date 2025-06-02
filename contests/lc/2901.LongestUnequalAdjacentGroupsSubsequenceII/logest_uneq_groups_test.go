package _2901_longest_uneq_groups_subsequence_ii

import (
	"testing"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		groups   []int
		expected int // только длина проверяем
	}{
		{
			name:     "Example 1",
			words:    []string{"bab", "dab", "cab"},
			groups:   []int{1, 2, 2},
			expected: 2, // ["bab", "dab"] или ["bab", "cab"]
		},
		{
			name:     "Single element",
			words:    []string{"abc"},
			groups:   []int{1},
			expected: 1,
		},
		{
			name:     "No valid hamming pairs",
			words:    []string{"abc", "xyz", "mno"},
			groups:   []int{1, 2, 1},
			expected: 1,
		},
		{
			name:     "Chain of 3",
			words:    []string{"abc", "bbc", "bcc", "ccc"},
			groups:   []int{1, 2, 1, 2},
			expected: 4, // full chain: abc -> bbc -> bcc -> ccc
		},
		{
			name:     "Same group block",
			words:    []string{"abc", "bbc", "acc"},
			groups:   []int{1, 1, 1},
			expected: 1, // все из одной группы, нельзя продолжать
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getWordsInLongestSubsequence(tc.words, tc.groups)
			if len(got) != tc.expected {
				t.Errorf("Expected length %d, got %d. Result: %v", tc.expected, len(got), got)
			}
		})
	}
}
