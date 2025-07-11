package meetingroomsiii

import (
	"testing"
)

func Test_mostBooked(t *testing.T) {
	type args struct {
		n        int
		meetings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				n:        2,
				meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			},
			want: 0,
		},
		{
			name: "test case 2",
			args: args{
				n:        3,
				meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostBooked(tt.args.n, tt.args.meetings); got != tt.want {
				t.Errorf("mostBooked() = %v, want %v", got, tt.want)
			}
		})
	}
}
