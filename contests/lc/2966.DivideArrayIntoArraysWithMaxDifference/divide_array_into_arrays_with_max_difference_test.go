package _2966_DivideArrayIntoArraysWithMaxDifference

import (
	"reflect"
	"testing"
)

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
				k:    2,
			},
			want: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 4, 2, 2, 5, 2},
				k:    2,
			},
			want: [][]int{},
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11},
				k:    14,
			},
			want: [][]int{{2, 2, 2}, {4, 5, 5}, {5, 5, 7}, {7, 8, 8}, {9, 9, 10}, {11, 12, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArray(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
