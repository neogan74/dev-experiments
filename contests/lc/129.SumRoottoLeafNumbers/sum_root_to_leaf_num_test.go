package _129_sum_root_to_leaf_num

import "testing"

func Test_sumNumbers(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 25, // 12 + 13
		},
		{
			name: "Example 2",
			root: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 9,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 0},
			},
			expected: 1026, // 495 + 491 + 40
		},
		{
			name:     "Empty",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single",
			root:     &TreeNode{Val: 9},
			expected: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sumNumbers(tc.root)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
