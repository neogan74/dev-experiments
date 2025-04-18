package main

import (
	"reflect"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	tests := []struct {
		score    []int
		expected []string
	}{
		{
			score:    []int{5, 4, 3, 2, 1},
			expected: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			score:    []int{10, 3, 8, 9, 4},
			expected: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			score:    []int{1},
			expected: []string{"Gold Medal"},
		},
	}

	for _, tt := range tests {
		result := findRelativeRanks(tt.score)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("findRelativeRanks(%v) = %v; want %v", tt.score, result, tt.expected)
		}
	}
}
