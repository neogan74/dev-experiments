package _169_CountOperationstoObtainZero

import "testing"

func TestCountOperations(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{
			name:     "Example 1: num1=2, num2=3",
			num1:     2,
			num2:     3,
			expected: 3,
		},
		{
			name:     "Example 2: num1=10, num2=10",
			num1:     10,
			num2:     10,
			expected: 1,
		},
		{
			name:     "Both zero",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "First zero",
			num1:     0,
			num2:     5,
			expected: 0,
		},
		{
			name:     "Second zero",
			num1:     5,
			num2:     0,
			expected: 0,
		},
		{
			name:     "num1 > num2",
			num1:     7,
			num2:     3,
			expected: 5,
		},
		{
			name:     "num2 > num1",
			num1:     3,
			num2:     7,
			expected: 5,
		},
		{
			name:     "Both are 1",
			num1:     1,
			num2:     1,
			expected: 1,
		},
		{
			name:     "Large numbers",
			num1:     100,
			num2:     10,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countOperations(tt.num1, tt.num2)
			if result != tt.expected {
				t.Errorf("countOperations(%d, %d) = %d; expected %d", tt.num1, tt.num2, result, tt.expected)
			}
		})
	}
}
