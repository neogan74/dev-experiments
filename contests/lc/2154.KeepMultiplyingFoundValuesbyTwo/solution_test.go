package keepmultiplyingfoundvaluesbytwo

import "testing"

func Test_findFinalValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums     []int
		original int
		want     int
	}{
		{
			name:     "Test Case 1",
			nums:     []int{5, 3, 6, 1, 12},
			original: 3,
			want:     24,
		},
		{
			name:     "Test Case 2",
			nums:     []int{2, 7, 9},
			original: 4,
			want:     4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFinalValue(tt.nums, tt.original)
			if got != tt.want {
				t.Errorf("findFinalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
