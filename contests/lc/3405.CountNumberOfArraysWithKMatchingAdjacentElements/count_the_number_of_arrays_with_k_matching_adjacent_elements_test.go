package _3405_CountNumberOfArraysWithKMatchingAdjacentElements

import "testing"

func Test_countGoodArrays(t *testing.T) {
	type args struct {
		n int
		m int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: 3,
				m: 2,
				k: 1,
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				n: 4,
				m: 2,
				k: 2,
			},
			want: 6,
		},
		{
			name: "Test 3",
			args: args{
				n: 5,
				m: 2,
				k: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodArrays(tt.args.n, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("countGoodArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
