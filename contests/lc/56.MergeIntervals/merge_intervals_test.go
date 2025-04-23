package _6_MergeIntervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			input:    [][]int{{1, 4}, {0, 4}},
			expected: [][]int{{0, 4}},
		},
		{
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			input:    [][]int{{1, 4}},
			expected: [][]int{{1, 4}},
		},
	}

	for _, tt := range tests {
		got := merge(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("merge(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
