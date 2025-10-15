package _539_FindSumofArrayProductofMagicalSequences

import "testing"

func Test_magicalSum(t *testing.T) {
	type args struct {
		m    int
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				m:    5,
				k:    5,
				nums: []int{1, 10, 100, 10000, 1_000_000},
			},
			want: 991600007,
		},
		{
			name: "test 2",
			args: args{
				m:    2,
				k:    2,
				nums: []int{5, 4, 3, 2, 1},
			},
			want: 170,
		},
		{
			name: "test 2",
			args: args{
				m:    1,
				k:    1,
				nums: []int{28},
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicalSum(tt.args.m, tt.args.k, tt.args.nums); got != tt.want {
				t.Errorf("magicalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
