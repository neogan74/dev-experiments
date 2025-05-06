package _25_reverse_node

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2, 1}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}}, // меньше чем k
	}

	for i, tc := range tests {
		head := sliceToList(tc.input)
		got := listToSlice(reverseKGroup(head, tc.k))
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("Test %d failed: expected %v, got %v", i+1, tc.expected, got)
		}
	}
}
