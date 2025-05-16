package _2094_finding_3_dig_even_num

import (
	"reflect"
	"testing"
)

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{

			digits: []int{2, 1, 3, 0},
			want:   []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
		{

			digits: []int{2, 2, 8, 8, 2},
			want:   []int{222, 228, 282, 288, 822, 828, 882},
		},
		{
			digits: []int{0, 0, 0},
			want:   []int{},
		},
		{
			digits: []int{1, 4, 7},
			want:   []int{174, 714}, // there is no even
		},
		{
			digits: []int{0, 1, 2},
			want:   []int{102, 120, 210},
		},
	}
	for _, tc := range tests {
		got := findNumbers(tc.digits)
		if got == nil {
			got = []int{}
		}
		if tc.want == nil {
			tc.want = []int{}
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("findNumbers() = %v, want %v", got, tc.want)
		}
	}
}
