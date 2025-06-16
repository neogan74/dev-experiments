package _2894_DivisibleandNon_divisibleSumsDifference

import "testing"

func Test_differenceOfSum(t *testing.T) {
	type args struct {
		n int64
		m int64
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "test 1",
			args: args{
				n: 10,
				m: 3,
			},
			wantAns: 19,
		},
		{
			name: "test 2",
			args: args{
				n: 5,
				m: 6,
			},
			wantAns: 15,
		},
		{
			name: "test 3",
			args: args{
				n: 5,
				m: 1,
			},
			wantAns: -15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := differenceOfSums(tt.args.n, tt.args.m); gotAns != tt.wantAns {
				t.Errorf("differenceOfSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
