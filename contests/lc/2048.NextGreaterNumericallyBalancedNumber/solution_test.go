package _048_NextGreaterNumericallyBalancedNumber

import "testing"

func Test_nextBeautifulNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: 1,
			},
			want: 22,
		},
		{
			name: "Test 2",
			args: args{
				n: 1000,
			},
			want: 1333,
		},
		{
			name: "Test 3",
			args: args{
				n: 3000,
			},
			want: 3133,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextBeautifulNumber(tt.args.n); got != tt.want {
				t.Errorf("nextBeautifulNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
