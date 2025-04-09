package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			input: [][]int{
				{5},
			},
			expected: [][]int{
				{5},
			},
		},
	}

	for _, tt := range tests {
		// Копируем матрицу, чтобы не мутировать оригинал в expected
		matrix := deepCopy(tt.input)
		rotate(matrix)
		if !reflect.DeepEqual(matrix, tt.expected) {
			t.Errorf("rotate(%v) = %v; want %v", tt.input, matrix, tt.expected)
		}
	}
}

func deepCopy(matrix [][]int) [][]int {
	copyMatrix := make([][]int, len(matrix))
	for i := range matrix {
		copyMatrix[i] = make([]int, len(matrix[i]))
		copy(copyMatrix[i], matrix[i])
	}
	return copyMatrix
}
