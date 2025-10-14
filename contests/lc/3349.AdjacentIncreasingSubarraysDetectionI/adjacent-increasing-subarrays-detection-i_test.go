package _349_AdjacentIncreasingSubarraysDetectionI

import "testing"

func Test_hasUniqueElements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7},
				k:    5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasUniqueElements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("hasUniqueElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
