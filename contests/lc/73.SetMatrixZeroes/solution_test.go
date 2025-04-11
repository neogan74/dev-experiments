package main

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, tt := range tests {
		inputCopy := deepCopy(tt.input)
		setZeroes(inputCopy)
		if !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("setZeroes(...) = %v; want %v", inputCopy, tt.expected)
		}
	}
}

func deepCopy(matrix [][]int) [][]int {
	copyMatrix := make([][]int, len(matrix))
	for i := range matrix {
		copyMatrix[i] = append([]int(nil), matrix[i]...)
	}
	return copyMatrix
}
