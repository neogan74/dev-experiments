package main

import "testing"

func intPtr(v int) *int {
	return &v
}

func buildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &TreeNode{Val: *v}
		}
	}
	for i := range nodes {
		if nodes[i] == nil {
			continue
		}
		li := 2*i + 1
		ri := 2*i + 2
		if li < len(nodes) {
			nodes[i].Left = nodes[li]
		}
		if ri < len(nodes) {
			nodes[i].Right = nodes[ri]
		}
	}
	return nodes[0]
}

func TestMaxLevelSumExamples(t *testing.T) {
	root := buildTree([]*int{
		intPtr(1), intPtr(7), intPtr(0), intPtr(7), intPtr(-8), nil, nil,
	})
	if got := maxLevelSum(root); got != 2 {
		t.Fatalf("example1: got %d, want 2", got)
	}

	root = buildTree([]*int{
		intPtr(989), nil, intPtr(10250), intPtr(98693), intPtr(-89388), nil, nil, nil, intPtr(-32127),
	})
	if got := maxLevelSum(root); got != 2 {
		t.Fatalf("example2: got %d, want 2", got)
	}
}

func TestMaxLevelSumSingleNode(t *testing.T) {
	root := buildTree([]*int{intPtr(-5)})
	if got := maxLevelSum(root); got != 1 {
		t.Fatalf("single node: got %d, want 1", got)
	}
}
