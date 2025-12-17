package binaryprefixdivisibleby5

import (
	"reflect"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []bool
	}{
		{
			name: "single zero",
			nums: []int{0},
			want: []bool{true},
		},
		{
			name: "example 1",
			nums: []int{0, 1, 1},
			want: []bool{true, false, false},
		},
		{
			name: "example 2",
			nums: []int{1, 1, 1},
			want: []bool{false, false, false},
		},
		{
			name: "alternating bits",
			nums: []int{1, 0, 1, 0, 1},
			want: []bool{false, false, true, true, false},
		},
		{
			name: "long run of ones",
			nums: []int{1, 1, 1, 1, 1, 1},
			want: []bool{false, false, false, true, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixesDivBy5(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("prefixesDivBy5(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
