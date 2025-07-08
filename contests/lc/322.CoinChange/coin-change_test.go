package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Example 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			name:     "Example 2",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Example 3",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.coins, tt.amount); got != tt.expected {
				t.Errorf("coinChange() = %v, want %v", got, tt.expected)
			}
		})
	}
}
