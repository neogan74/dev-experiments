package _461_CheckIfDigitsAreEqualinStringAfterOperationsI

import "testing"

func Test_hasSameDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "3902",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: "34789",
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				s: "1234567890",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSameDigits(tt.args.s); got != tt.want {
				t.Errorf("hasSameDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
