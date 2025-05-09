package _1920_build_array_from_permutation

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{0, 2, 1, 5, 3, 4}, expected: []int{0, 1, 2, 4, 5, 3}},
		{input: []int{5, 0, 1, 2, 3, 4}, expected: []int{4, 5, 0, 1, 2, 3}},
		{input: []int{1, 0}, expected: []int{0, 1}},
		{input: []int{0}, expected: []int{0}},
	}

	for i, tc := range tests {
		got := buildArray(tc.input)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("Test %d failed: expected %v, got %v", i+1, tc.expected, got)
		}
	}
}
