package main

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "increasing", input: []int{1, 2, 3}, expected: []int{1, 3, 2}},
		{name: "decreasing wraps", input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
		{name: "with duplicates", input: []int{1, 1, 5}, expected: []int{1, 5, 1}},
		{name: "middle pivot", input: []int{1, 3, 2}, expected: []int{2, 1, 3}},
		{name: "single element", input: []int{7}, expected: []int{7}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.input...)
			nextPermutation(nums)
			if !reflect.DeepEqual(nums, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, nums)
			}
		})
	}
}
