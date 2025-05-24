package _236_lowest_common_ancestor_of_bt

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	//    3
	//  /   \
	//  5   1
	// / \  /\
	//6  2 0  8
	//  / \
	// 7   4
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n6 := &TreeNode{Val: 6}
	n0 := &TreeNode{Val: 0}
	n8 := &TreeNode{Val: 8}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}
	root := &TreeNode{Val: 3, Left: n5, Right: n1}

	tests := []struct {
		p, q     *TreeNode
		expected int
	}{
		{n5, n1, 3},
		{n6, n4, 5},
		{n7, n8, 3},
		{n4, n5, 5},
	}
	for _, tc := range tests {
		got := lowestCommonAncestor(root, tc.p, tc.q)
		if got == nil || got.Val != tc.expected {
			t.Errorf("lowestCommonAncestor() = %v, want %v", got, tc.expected)
		}

	}
}
