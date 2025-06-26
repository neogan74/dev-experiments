package longestbinarysubsequencelessthanorequaltok

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example 1",
			args: args{
				s: "1001010",
				k: 5,
			},
			wantAns: 5,
		},
		{
			name: "Example 2",
			args: args{
				s: "00101001",
				k: 1,
			},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSubsequence(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("longestSubsequence() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
