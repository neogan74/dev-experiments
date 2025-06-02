package _2894_DivisibleandNon_divisibleSumsDifference

import "testing"

func Test_differenceOfSum(t *testing.T) {
	type args struct {
		nums []int
		m    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				m:    2,
			},
			wantAns: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := differenceOfSum(tt.args.nums, tt.args.m); gotAns != tt.wantAns {
				t.Errorf("differenceOfSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
