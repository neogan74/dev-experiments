package _488_AvoidFloodinTheCity

import (
	"reflect"
	"testing"
)

func TestAvoidFlood(t *testing.T) {
	type args struct {
		rains []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				rains: []int{
					1, 2, 3, 4,
				},
			},
			want: []int{-1, -1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AvoidFlood(tt.args.rains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AvoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
