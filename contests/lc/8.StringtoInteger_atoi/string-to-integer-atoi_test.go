package stringtointegeratoi

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "Test 1",
			s:    "42",
			want: 42,
		},
		{
			name: "Test 2",
			s:    "   -42",
			want: -42,
		},
		{
			name: "Test 3",
			s:    "4193 with words",
			want: 4193,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.s)
			if got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
