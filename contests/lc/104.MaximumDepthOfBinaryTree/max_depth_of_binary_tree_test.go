package _104_maximum_depth_of_binary_tree

import "testing"

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		tree     *TreeNode
		expected int
	}{
		{
			tree:     nil,
			expected: 0,
		},
		{
			tree:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 3,
		},
		{
			tree: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: 3,
		},
	}

	for i, tc := range tests {
		got := maxDepth(tc.tree)
		if got != tc.expected {
			t.Errorf("Test %d failed: expected %d, got %d", i+1, tc.expected, got)
		}
	}
}
