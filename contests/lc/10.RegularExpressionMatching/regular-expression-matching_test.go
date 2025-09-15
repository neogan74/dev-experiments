package regularexpressionmatching

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "Example 2",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "Example 3",
			s:    "ab",
			p:    ".*",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMatch(tt.s, tt.p)
			if got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
