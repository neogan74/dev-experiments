package _52_n_queens2

import "testing"

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Test 1",
			args: args{
				n: 4,
			},
			wantAns: 2,
		},
		{
			name: "Test 2",
			args: args{
				n: 1,
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := totalNQueens(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("totalNQueens() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
