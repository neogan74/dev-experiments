package main

import (
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	tests := []struct {
		board    [][]int
		expected [][]int
	}{
		{
			board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			board: [][]int{
				{1, 1},
				{1, 0},
			},
			expected: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for _, tt := range tests {
		input := deepCopy(tt.board)
		gameOfLife(input)
		if !reflect.DeepEqual(input, tt.expected) {
			t.Errorf("gameOfLife(...) = %v; want %v", input, tt.expected)
		}
	}
}

func deepCopy(grid [][]int) [][]int {
	result := make([][]int, len(grid))
	for i := range grid {
		result[i] = append([]int(nil), grid[i]...)
	}
	return result
}
