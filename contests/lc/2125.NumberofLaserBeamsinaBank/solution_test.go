package numberoflaserbeamsinabank

import "testing"

func Test_numberOfBeams(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank []string
		want int
	}{
		{
			name: "Example 1",
			bank: []string{"011001", "000000", "010100", "001000"},
			want: 8,
		},
		{
			name: "Example 2",
			bank: []string{"000", "111", "000"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfBeams(tt.bank)
			if got != tt.want {
				t.Errorf("numberOfBeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
