package maximumfruitsharvestedafteratmostksteps

import "testing"

func Test_maxTotalFruits(t *testing.T) {
	type args struct {
		fruits   [][]int
		startPos int
		k        int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Test 1",
			args: args{
				fruits:   [][]int{{2, 8}, {6, 3}, {8, 6}},
				startPos: 5,
				k:        4,
			},
			wantAns: 9,
		},
		{
			name: "Test 2",
			args: args{
				fruits:   [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}},
				startPos: 5,
				k:        4,
			},
			wantAns: 14,
		},
		{
			name: "Test 3",
			args: args{
				fruits:   [][]int{{0, 3}, {6, 4}, {8, 5}},
				startPos: 3,
				k:        2,
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxTotalFruits(tt.args.fruits, tt.args.startPos, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxTotalFruits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
