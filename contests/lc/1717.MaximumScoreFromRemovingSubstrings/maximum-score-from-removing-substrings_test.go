package maximumscorefromremovingsubstrings

import "testing"

func Test_maximumGain(t *testing.T) {
	type args struct {
		s string
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "test1",
			args: args{
				s: "cdbcbbaaabab",
				x: 4,
				y: 5,
			},
			wantAns: 19,
		},
		{
			name: "test2",
			args: args{
				s: "aabbaaxybbaabb",
				x: 5,
				y: 4,
			},
			wantAns: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumGain(tt.args.s, tt.args.x, tt.args.y); gotAns != tt.wantAns {
				t.Errorf("maximumGain() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
