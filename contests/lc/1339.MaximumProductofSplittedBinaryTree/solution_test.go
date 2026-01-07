package main

import "testing"

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

func TestMaxProductExamples(t *testing.T) {
	tests := []struct {
		name   string
		input  []*int
		output int
	}{
		{
			name:   "example1",
			input:  []*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6)},
			output: 110,
		},
		{
			name:   "example2",
			input:  []*int{ptr(1), nil, ptr(2), ptr(3), ptr(4), nil, nil, ptr(5), ptr(6)},
			output: 90,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			root := buildTree(test.input)
			if got := maxProduct(root); got != test.output {
				t.Fatalf("maxProduct() = %d, want %d", got, test.output)
			}
		})
	}
}
