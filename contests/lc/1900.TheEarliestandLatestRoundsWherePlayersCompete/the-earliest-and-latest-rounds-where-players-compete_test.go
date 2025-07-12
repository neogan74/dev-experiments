package theearliestandlatestroundswhereplayerscompete

import (
	"reflect"
	"testing"
)

func Test_earliestAndLatest(t *testing.T) {
	type args struct {
		n            int
		firstPlayer  int
		secondPlayer int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				n:            11,
				firstPlayer:  2,
				secondPlayer: 4,
			},
			want: []int{3, 4},
		},
		{
			name: "Example 2",
			args: args{
				n:            5,
				firstPlayer:  1,
				secondPlayer: 5,
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestAndLatest(tt.args.n, tt.args.firstPlayer, tt.args.secondPlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("earliestAndLatest() = %v, want %v", got, tt.want)
			}
		})
	}
}
