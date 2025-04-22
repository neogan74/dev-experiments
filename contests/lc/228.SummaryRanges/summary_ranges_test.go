package _28_SummaryRanges

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []string
	}{
		{[]int{}, []string{}},
		{[]int{1}, []string{"1"}},
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{-3, -2, -1, 1, 3}, []string{"-3->-1", "1", "3"}},
		{[]int{5, 6, 7, 8}, []string{"5->8"}},
	}

	for _, tt := range tests {
		got := summaryRanges(tt.nums)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("summaryRanges(%v) = %v; want %v", tt.nums, got, tt.expected)
		}
	}
}
