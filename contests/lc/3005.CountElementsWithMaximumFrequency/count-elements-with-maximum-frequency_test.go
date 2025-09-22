package countelementswithmaximumfrequency

import "testing"

func Test_maxFrequencyElements(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Test Case 1",
			nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			want: 3,
		},
		{
			name: "Test Case 2",
			nums: []int{1, 1, 2, 2, 3, 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFrequencyElements(tt.nums)
			if got != tt.want {
				t.Errorf("maxFrequencyElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
