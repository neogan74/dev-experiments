package _100_same_tree

import "testing"

func Test_isSameTree(t *testing.T) {
	tests := []struct {
		name     string
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			name:     "Both nil",
			p:        nil,
			q:        nil,
			expected: true,
		},
		{
			name: "Same structure and values",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			name: "Different structure",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name: "Different values",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
			},
			expected: false,
		},
	}

	for i, tc := range tests {
		got := isSameTree(tc.p, tc.q)
		if got != tc.expected {
			t.Errorf("Test %d (%s) failed: expected %v, got %v", i+1, tc.name, tc.expected, got)
		}
	}
}
