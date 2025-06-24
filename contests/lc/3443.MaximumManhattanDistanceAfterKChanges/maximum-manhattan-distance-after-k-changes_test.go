package _3443_maximummanhattandistanceafterkchanges

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		s string
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
				s: "NWSE",
				k: 1,
			},
			want: 3,
		},
		{
			name: "Test 1",
			args: args{
				s: "NSWWEW",
				k: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
