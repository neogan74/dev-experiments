package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"testing"
)

func equalTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return equalTree(a.Left, b.Left) && equalTree(a.Right, b.Right)
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		inorder   []int
		postorder []int
		want      *TreeNode
	}{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			inorder:   []int{2, 1},
			postorder: []int{2, 1},
			want: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2},
			},
		},
		{
			inorder:   []int{},
			postorder: []int{},
			want:      nil,
		},
	}

	for i, tc := range tests {
		got := buildTree(tc.inorder, tc.postorder)
		if !equalTree(got, tc.want) {
			t.Errorf("Test %d failed:\n  inorder=%v\n  postorder=%v\n  expected=%+v\n  got=%+v", i+1, tc.inorder, tc.postorder, tc.want, got)
		}
	}
}
