package _399_CountLargestGroup

import "testing"

func Test_countLargestGroup(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example 1",
			args: args{
				n: 13,
			},
			wantAns: 4,
		},
		{
			name: "Example 2",
			args: args{
				n: 2,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countLargestGroup(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("countLargestGroup() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
