package _75_sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element 0",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "single element 1",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "single element 2",
			input:    []int{2},
			expected: []int{2},
		},
		{
			name:     "already sorted",
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "reverse sorted",
			input:    []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "mixed order",
			input:    []int{2, 0, 1, 1, 0, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "all 0s",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "all 1s",
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "all 2s",
			input:    []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
		{
			name:     "random order 1",
			input:    []int{1, 2, 0, 1, 2, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "random order 2",
			input:    []int{2, 1, 0, 0, 1, 2, 1, 0},
			expected: []int{0, 0, 0, 1, 1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to preserve it for error messages
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			sortColors(tt.input)

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("sortColors(%v) = %v, want %v", inputCopy, tt.input, tt.expected)
			}
		})
	}
}
