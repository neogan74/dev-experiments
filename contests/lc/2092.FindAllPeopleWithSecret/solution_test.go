package _092_FindAllPeopleWithSecret

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAllPeople(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		meetings    [][]int
		firstPerson int
		expected    []int
	}{
		{
			name:        "example 1",
			n:           6,
			firstPerson: 1,
			meetings: [][]int{
				{1, 2, 5},
				{2, 3, 5},
				{3, 4, 5},
				{1, 5, 5},
			},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:        "example 2",
			n:           4,
			firstPerson: 3,
			meetings: [][]int{
				{2, 3, 7},
				{1, 3, 7},
				{0, 2, 7},
				{0, 3, 7},
			},
			expected: []int{0, 1, 2, 3},
		},
		{
			name:        "no meetings",
			n:           2,
			firstPerson: 1,
			meetings:    [][]int{},
			expected:    []int{0, 1},
		},
		{
			name:        "secret not spread",
			n:           5,
			firstPerson: 4,
			meetings: [][]int{
				{1, 2, 10},
				{2, 3, 10},
				{1, 3, 11},
			},
			expected: []int{0, 4},
		},
		{
			name:        "single meeting with spread",
			n:           3,
			firstPerson: 0,
			meetings: [][]int{
				{0, 1, 5},
				{1, 2, 5},
			},
			expected: []int{0, 1, 2},
		},
		{
			name:        "disconnected components",
			n:           6,
			firstPerson: 1,
			meetings: [][]int{
				{1, 2, 5},
				{3, 4, 5},
				{2, 5, 5},
				{0, 1, 5},
			},
			expected: []int{0, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAllPeople(tt.n, tt.meetings, tt.firstPerson)

			// Проверяем, что результаты совпадают с учётом порядка
			sort.Ints(result)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findAllPeople() = %v, want %v", result, tt.expected)
			}
		})
	}
}
