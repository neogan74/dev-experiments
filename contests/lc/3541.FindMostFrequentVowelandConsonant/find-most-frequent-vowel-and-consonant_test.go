package findmostfrequentvowelandconsonant

import "testing"

func Test_maxFreqSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "leetcode",
			want: 4,
		},
		{
			name: "Example 2",
			s:    "abbcccddddeeeeedcba",
			want: 10,
		},
		{
			name: "Example 3",
			s:    "successes",
			want: 6,
		},
		{
			name: "Custom 1",
			s:    "aeiaeia",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFreqSum(tt.s)
			if got != tt.want {
				t.Errorf("maxFreqSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
