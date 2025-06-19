package _39_CombinationSum

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Test 1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			wantAns: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "Test 2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			wantAns: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "Test 3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
