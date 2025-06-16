package _3170_lex_min_adter_rwmovw

import "testing"

func TestClearStars(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{s: "aaba*", expected: "aab"},
		{s: "abc", expected: "abc"},
	}
	for _, tt := range tests {
		result := clearStars(tt.s)
		if result != tt.expected {
			t.Errorf("Something wrong got: %v, expected: %v", result, tt.expected)
		}
	}
}
