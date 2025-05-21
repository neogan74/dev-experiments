package _24_BinaryTreeMaximumPathSum

import "testing"

func Test_maxPathSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{Val: -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				}},
			expected: 42,
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "All negative",
			root: &TreeNode{Val: -3,
				Left: &TreeNode{Val: -2,
					Left: &TreeNode{Val: -1},
				},
			},
			expected: -1,
		},
		{
			name: "Mixed values",
			root: &TreeNode{Val: 2,
				Left: &TreeNode{Val: -1},
			},
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxPathSum(tc.root)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
