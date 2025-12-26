package combinationsumii

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Test case 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{2, 6},
				{1, 7},
				{1, 2, 5},
				{1, 1, 6},
			},
		},
		{
			name:       "Test case 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected: [][]int{
				{5},
				{1, 2, 2},
			},
		},
		{
			name:       "Test case 3 (no combinations)",
			candidates: []int{2, 3, 6, 7},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "Test case 4 (empty candidates)",
			candidates: []int{},
			target:     5,
			expected:   [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum2(tt.candidates, tt.target)

			// Сортировка для корректного сравнения
			sort.Slice(result, func(i, j int) bool {
				return !equalSlice(result[i], result[j])
			})

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("combinationSum2(%v, %d) = %v; expected %v", tt.candidates, tt.target, result, tt.expected)
			}
		})
	}
}

// Вспомогательная функция для сортировки результата
func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
