package rearrangingfruits

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		basket1 []int
		basket2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "test 1",
			args: args{
				basket1: []int{4, 2, 2, 2},
				basket2: []int{1, 4, 1, 2},
			},
			wantAns: 1,
		},
		{
			name: "test 2",
			args: args{
				basket1: []int{2, 3, 4, 1},
				basket2: []int{3, 2, 5, 1},
			},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minCost(tt.args.basket1, tt.args.basket2); gotAns != tt.wantAns {
				t.Errorf("minCost() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
