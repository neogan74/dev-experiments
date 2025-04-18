package main

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"abc", "one two three", true},
		{"ab", "one one", false},
	}

	for _, tt := range tests {
		result := wordPattern(tt.pattern, tt.s)
		if result != tt.expected {
			t.Errorf("wordPattern(%q, %q) = %v; want %v",
				tt.pattern, tt.s, result, tt.expected)
		}
	}
}
