package __112_path_sum

import "testing"

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expected  bool
	}{
		{
			name: "Example 1",
			root: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 4,
					Left: &TreeNode{Val: 11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					}},
				Right: &TreeNode{Val: 8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{Val: 4,
						Right: &TreeNode{Val: 1},
					}},
			},
			targetSum: 22,
			expected:  true,
		},
		{
			name:      "Empty tree",
			root:      nil,
			targetSum: 0,
			expected:  false,
		},
		{
			name:      "Single-node match",
			root:      &TreeNode{Val: 7},
			targetSum: 7,
			expected:  true,
		},
		{
			name:      "Single-node mismatch",
			root:      &TreeNode{Val: 1},
			targetSum: 2,
			expected:  false,
		},
		{
			name: "Path exists but ends not at leaf",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2},
			},
			targetSum: 1,
			expected:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := hasPathSum(tc.root, tc.targetSum)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
