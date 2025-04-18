package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			input: []string{""},
			expected: [][]string{
				{""},
			},
		},
		{
			input: []string{"a"},
			expected: [][]string{
				{"a"},
			},
		},
	}

	for _, tt := range tests {
		got := groupAnagrams(tt.input)

		// Нормализуем порядок групп и порядок внутри групп
		sortGroup := func(groups [][]string) {
			for _, g := range groups {
				sort.Strings(g)
			}
			sort.Slice(groups, func(i, j int) bool {
				return strings.Join(groups[i], "") < strings.Join(groups[j], "")
			})
		}

		sortGroup(tt.expected)
		sortGroup(got)

		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("groupAnagrams(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
