package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3}, // "abc"
		{"bbbbb", 1},    // "b"
		{"pwwkew", 3},   // "wke"
		{"", 0},         // пустая строка
		{"abcdef", 6},   // все уникальные
		{"abba", 2},     // "ab" или "ba"
		{"a", 1},        // один символ
		{"dvdf", 3},     // "vdf"
		{"anviaj", 5},   // "nviaj"
		{"一二三二一", 3},    // Unicode: "一二三"
	}

	for _, tt := range tests {
		result := lengthOfLongestSubstring(tt.input)
		if result != tt.expected {
			t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}
