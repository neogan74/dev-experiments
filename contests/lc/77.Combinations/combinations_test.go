package _77_Combinations

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Test 1",
			args: args{
				n: 4,
				k: 2,
			},
			wantAns: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "Test 2",
			args: args{
				n: 1,
				k: 1,
			},
			wantAns: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combine() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
