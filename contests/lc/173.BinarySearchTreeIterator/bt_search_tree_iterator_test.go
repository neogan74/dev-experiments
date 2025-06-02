package _173_binary_search_tree_iterator

import "testing"

func TestBSTIterator(t *testing.T) {
	// BST:
	//       7
	//      / \
	//     3   15
	//         / \
	//        9  20
	root := &TreeNode{Val: 7,
		Left: &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20},
		},
	}

	it := Constructor(root)
	var result []int
	for it.HasNext() {
		result = append(result, it.Next())
	}

	expected := []int{3, 7, 9, 15, 20}
	if len(result) != len(expected) {
		t.Fatalf("Length mismatch: got %v, expected %v", result, expected)
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("At index %d: got %d, expected %d", i, result[i], expected[i])
		}
	}
}
