package main

import "testing"

func TestCheckPerfectNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{28, true},
		{6, true},
		{496, true},
		{8128, true},
		{2, false},
		{12, false},
		{1, false},
		{33550336, true}, // known perfect number
	}

	for _, tt := range tests {
		result := checkPerfectNumber(tt.input)
		if result != tt.expected {
			t.Errorf("checkPerfectNumber(%d) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
