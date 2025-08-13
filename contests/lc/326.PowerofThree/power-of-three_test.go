package powerofthree

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with 27",
			args: args{n: 27},
			want: true,
		},
		{
			name: "Test with 0",
			args: args{n: 0},
			want: false,
		},
		{
			name: "Test with -1",
			args: args{n: -1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
