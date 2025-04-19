package _8_CountAndSay

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
		{9, "31131211131221"},
		{10, "13211311123113112211"},
	}

	for _, tt := range tests {
		got := countAndSay(tt.n)
		if got != tt.expected {
			t.Errorf("countAndSay(%d) = %q; want %q", tt.n, got, tt.expected)
		}
	}
}
