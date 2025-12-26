package deletecolumnstomakesorted

import (
	"testing"
)

func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected int
	}{
		{
			name:     "Test 1",
			strs:     []string{"cba", "daf", "ghi"},
			expected: 1,
		},
		{
			name:     "Test 2",
			strs:     []string{"abc", "bce", "cae"},
			expected: 1,
		},
		{
			name:     "Test 3",
			strs:     []string{"abc", "acb", "bac"},
			expected: 2,
		},
		{
			name:     "Test 4",
			strs:     []string{"a", "b", "c"},
			expected: 0,
		},
		{
			name:     "Test 5",
			strs:     []string{"abcde"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDeletionSize(tt.strs)
			if result != tt.expected {
				t.Errorf("minDeletionSize(%v) = %d, want %d", tt.strs, result, tt.expected)
			}
		})
	}
}
