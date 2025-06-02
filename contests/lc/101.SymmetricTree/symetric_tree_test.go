package _101_symetric_tree

import "testing"

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name: "Symmetric tree",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3}},
			},
			expected: true,
		},
		{
			name: "Not symmetric",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2,
					Right: &TreeNode{Val: 3}},
			},
			expected: false,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "Asymmetric values",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isSymmetric(tc.root)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
