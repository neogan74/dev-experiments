package _397_MaximumNumberofDistinctElementsAfterOperations

import "testing"

func Test_maxDistinctElements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 2, 3, 3, 4},
				k:    2,
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{4, 4, 4, 4},
				k:    1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistinctElements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxDistinctElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
