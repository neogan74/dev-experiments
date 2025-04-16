package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t     string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "a", true},
		{"abc", "cba", true},
		{"abc", "abcd", false},
		{"", "", true},
	}

	for _, tt := range tests {
		result := isAnagram(tt.s, tt.t)
		if result != tt.expected {
			t.Errorf("isAnagram(%q, %q) = %v; want %v", tt.s, tt.t, result, tt.expected)
		}
	}
}
