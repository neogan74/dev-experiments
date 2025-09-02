package findthenumberofwaystoplacepeoplei

import "testing"

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "test 1",
			args:    args{points: [][]int{{1, 1}, {2, 2}, {3, 3}}},
			wantAns: 0,
		},
		{
			name:    "test 2",
			args:    args{points: [][]int{{6, 2}, {4, 4}, {2, 6}}},
			wantAns: 2,
		},
		{
			name:    "test 3",
			args:    args{points: [][]int{{3, 1}, {1, 3}, {1, 1}}},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfPairs(tt.args.points); gotAns != tt.wantAns {
				t.Errorf("numberOfPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
