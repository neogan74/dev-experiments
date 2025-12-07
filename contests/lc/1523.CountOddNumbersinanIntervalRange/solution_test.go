package countoddnumbersinanintervalrange

import "testing"

func Test_countOdds(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		low  int
		high int
		want int
	}{
		{
			name: "Test Case 1",
			low:  3,
			high: 7,
			want: 3,
		},
		{
			name: "Test Case 2",
			low:  8,
			high: 10,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countOdds(tt.low, tt.high)
			if got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"2", func(t *testing.T) {
			got := countOdds2(tt.low, tt.high)
			if got != tt.want {
				t.Errorf("countOdds2() = %v, want %v", got, tt.want)
			}
		})
	}
}
