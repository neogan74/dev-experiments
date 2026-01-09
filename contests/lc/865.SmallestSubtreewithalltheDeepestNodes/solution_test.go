package smallestsubtreewithallthedeepestnodes

import (
	"reflect"
	"testing"
)

func ptr(v int) *int {
	return &v
}

func buildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	index := 1

	for index < len(values) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if index < len(values) {
			if values[index] != nil {
				node.Left = &TreeNode{Val: *values[index]}
				queue = append(queue, node.Left)
			}
			index++
		}

		if index < len(values) {
			if values[index] != nil {
				node.Right = &TreeNode{Val: *values[index]}
				queue = append(queue, node.Right)
			}
			index++
		}
	}

	return root
}

func treeToSlice(root *TreeNode) []any {
	if root == nil {
		return nil
	}

	result := make([]any, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func TestSubtreeWithAllDeepest(t *testing.T) {
	tests := []struct {
		name     string
		input    []*int
		expected []any
	}{
		{
			name:     "example1",
			input:    []*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(0), ptr(8), nil, nil, ptr(7), ptr(4)},
			expected: []any{2, 7, 4},
		},
		{
			name:     "example2",
			input:    []*int{ptr(1)},
			expected: []any{1},
		},
		{
			name:     "example3",
			input:    []*int{ptr(0), ptr(1), ptr(3), nil, ptr(2)},
			expected: []any{2},
		},
		{
			name:     "two_deepest_siblings",
			input:    []*int{ptr(1), ptr(2), ptr(3)},
			expected: []any{1, 2, 3},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			root := buildTree(test.input)
			got := subtreeWithAllDeepest(root)
			gotSlice := treeToSlice(got)
			if !reflect.DeepEqual(gotSlice, test.expected) {
				t.Fatalf("subtreeWithAllDeepest() = %v, want %v", gotSlice, test.expected)
			}
		})
	}
}
