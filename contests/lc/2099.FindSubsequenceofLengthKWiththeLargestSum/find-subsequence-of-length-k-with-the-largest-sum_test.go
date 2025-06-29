package findsubsequenceoflengthkwiththelargestsum

import (
	"reflect"
	"testing"
)

func Test_maxSubsequence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{2, 1, 3, 3},
				k:    2,
			},
			want: []int{3, 3},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{-1, -2, 3, 4},
				k:    3,
			},
			want: []int{-1, 3, 4},
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{3, 4, 3, 3},
				k:    2,
			},
			want: []int{4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsequence(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
