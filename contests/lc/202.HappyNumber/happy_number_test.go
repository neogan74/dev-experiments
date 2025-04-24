package _02_HappyNumber

import "testing"

func Test_sumOfSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{
				n: 19,
			},
			want: 82,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSquares(tt.args.n); got != tt.want {
				t.Errorf("sumOfSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{
				n: 19,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
