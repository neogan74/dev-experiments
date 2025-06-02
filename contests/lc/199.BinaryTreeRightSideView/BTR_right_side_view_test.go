package _199_binary_tree_right_side_view

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	//     1
	//    / \
	//   2   3
	//    \   \
	//     5   4
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}

	expected := []int{1, 3, 4}
	result := rightSideView(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestRightSideViewSingleNode(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}
	result := rightSideView(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
