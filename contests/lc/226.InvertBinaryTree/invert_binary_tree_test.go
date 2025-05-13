package _226_invert_binary_tree

import (
	"reflect"
	"testing"
)

func buildTreeFromLevelOrder(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	for i := 0; i < len(vals); i++ {
		if nodes[i] == nil {
			continue
		}
		leftIdx, rightIdx := 2*i+1, 2*i+2
		if leftIdx < len(vals) {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < len(vals) {
			nodes[i].Right = nodes[rightIdx]
		}
	}
	return nodes[0]
}

func treeEqual(a, b *TreeNode) bool {
	return reflect.DeepEqual(a, b)
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		input    []any
		expected []any
	}{
		{
			input:    []any{4, 2, 7, 1, 3, 6, 9},
			expected: []any{4, 7, 2, 9, 6, 3, 1},
		},
		{
			input:    []any{2, 1, 3},
			expected: []any{2, 3, 1},
		},
		{
			input:    []any{},
			expected: []any{},
		},
	}

	for i, tc := range tests {
		inTree := buildTreeFromLevelOrder(tc.input)
		expectedTree := buildTreeFromLevelOrder(tc.expected)
		got := invertTree(inTree)
		if !treeEqual(got, expectedTree) {
			t.Errorf("Test %d failed:\n  input:    %v\n  expected: %v\n  got:      %v", i+1, tc.input, tc.expected, got)
		}
	}
}
