package sortvowelsinastring

import (
	"testing"
)

func TestSortVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Leetcode example",
			input:    "lEetcOde",
			expected: "lEOtcede",
		},
		{
			name:     "All vowels lowercase",
			input:    "aeiou",
			expected: "aeiou",
		},
		{
			name:     "All vowels uppercase",
			input:    "AEIOU",
			expected: "AEIOU",
		},
		{
			name:     "No vowels",
			input:    "BCDFG",
			expected: "BCDFG",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Just set of words",
			input:    "lYmpH",
			expected: "lYmpH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortVowels(tt.input)
			if result != tt.expected {
				t.Errorf("sortVowels(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
