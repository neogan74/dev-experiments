package fractiontorecurringdecimal

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		numerator   int
		denominator int
		want        string
	}{
		{
			name:        "Example 1",
			numerator:   1,
			denominator: 2,
			want:        "0.5",
		},
		{
			name:        "Test 2",
			numerator:   2,
			denominator: 1,
			want:        "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fractionToDecimal(tt.numerator, tt.denominator)
			if got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
