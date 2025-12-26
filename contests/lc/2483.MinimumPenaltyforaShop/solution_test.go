package minimumpenaltyforashop

import "testing"

func Test_bestClosingTime(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		customers string
		want      int
	}{
		{
			name:      "Test Case 1",
			customers: "YYNY",
			want:      2,
		},
		{
			name:      "Test Case 2",
			customers: "NNNNN",
			want:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bestClosingTime(tt.customers)
			if got != tt.want {
				t.Errorf("bestClosingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
