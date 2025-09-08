package maximumareaoflongestdiagonalrectangle

import "testing"

func Test_areaOfMaxDiagonal(t *testing.T) {
	type args struct {
		dimensions [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Test 1",
			args: args{
				dimensions: [][]int{{9, 3}, {8, 6}},
			},
			wantAns: 48,
		},
		{
			name: "Test 2",
			args: args{
				dimensions: [][]int{{3, 4}, {4, 3}},
			},
			wantAns: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := areaOfMaxDiagonal(tt.args.dimensions); gotAns != tt.wantAns {
				t.Errorf("areaOfMaxDiagonal() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
