package utils

import (
	"reflect"
	"testing"
)

func treeToSlice(root *TreeNode) []interface{} {
	var result []interface{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	// trim trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]interface{} // возможны разные корректные деревья
	}{
		{
			nums: []int{-10, -3, 0, 5, 9},
			expected: [][]interface{}{
				{0, -10, 5, nil, -3, nil, 9},
				{0, -3, 9, -10, nil, 5},
			},
		},
		{
			nums: []int{1, 3},
			expected: [][]interface{}{
				{1, nil, 3},
				{3, 1},
			},
		},
	}

	for _, tt := range tests {
		result := treeToSlice(sortedArrayToBST(tt.nums))
		matched := false
		for _, exp := range tt.expected {
			if reflect.DeepEqual(result, exp) {
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("sortedArrayToBST(%v) = %v; expected one of %v", tt.nums, result, tt.expected)
		}
	}
}
