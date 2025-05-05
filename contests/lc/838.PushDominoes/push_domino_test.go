package _838_push_domino

import "testing"

func Test_pushDominoes(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    ".L.R...LR..L..",
			expected: "LL.RR.LLRRLL..",
		},
		{
			input:    "RR.L",
			expected: "RR.L",
		},
		{
			input:    "R....L",
			expected: "RRRLLL",
		},
		{
			input:    "...",
			expected: "...",
		},
		{
			input:    "L.....R",
			expected: "L.....R",
		},
		{
			input:    "R.R.L",
			expected: "RRR.L",
		},
	}

	for i, test := range tests {
		got := pushDominoes(test.input)
		if got != test.expected {
			t.Errorf("Test %d failed: input=%q, expected=%q, got=%q", i+1, test.input, test.expected, got)
		}
	}
}
