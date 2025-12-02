package findxsumofallklongsubarraysii

import (
	"reflect"
	"testing"
)

func Test_findXSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		x    int
		want []int64
	}{
		{
			name: "example 1",
			nums: []int{2, 4, 3, 5, 1},
			k:    3,
			x:    2,
			want: []int64{7, 9, 8},
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			x:    1,
			want: []int64{2, 3, 4, 5},
		},
		{
			name: "single element window",
			nums: []int{1, 2, 3},
			k:    1,
			x:    1,
			want: []int64{1, 2, 3},
		},
		{
			name: "equal elements",
			nums: []int{1, 1, 1, 1},
			k:    2,
			x:    1,
			want: []int64{2, 2, 2},
		},
		{
			name: "example 3",
			nums: []int{1, 1, 2, 2, 3, 4, 2, 3},
			k:    4,
			x:    2,
			want: []int64{6, 7, 8, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findXSum(tt.nums, tt.k, tt.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findXSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
