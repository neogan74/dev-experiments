package constructquadtree

import (
	"testing"
)

func Test_construct(t *testing.T) {
	grid := [][]int{
		{1, 1},
		{1, 1},
	}
	tree := construct(grid)
	if tree == nil || !tree.IsLeaf || !tree.Val {
		t.Errorf("Expected a leaf node with Val=true, got %+v", tree)
	}

	grid = [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
	}
	tree = construct(grid)
	if tree.IsLeaf {
		t.Errorf("Expected non-leaf root for mixed region")
	}

	grid = [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	tree = construct(grid)
	if tree.IsLeaf {
		t.Errorf("Expected non-leaf root for mixed region")
	}
}
