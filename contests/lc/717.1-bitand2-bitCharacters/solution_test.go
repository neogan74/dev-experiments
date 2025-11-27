package bitand2bitcharacters

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bits []int
		want bool
	}{
		{
			name: "Test Case 1",
			bits: []int{1, 0, 0},
			want: true,
		},
		{
			name: "Test Case 2",
			bits: []int{1, 1, 1, 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOneBitCharacter(tt.bits)
			if got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
