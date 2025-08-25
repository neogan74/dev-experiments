package four_sum

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:   "Example 2",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
