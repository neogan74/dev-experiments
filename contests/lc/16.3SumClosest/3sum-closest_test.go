package sumclosest

import "testing"

func Test_threeSumClosest(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Test 1",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			name:   "Test 2",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
