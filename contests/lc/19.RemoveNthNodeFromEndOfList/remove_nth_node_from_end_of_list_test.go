package _19_remove_nth_node_from_end_of_list

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

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input    []int
		n        int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, n: 2, expected: []int{1, 2, 3, 5}},
		{input: []int{1}, n: 1, expected: []int{}},
		{input: []int{1, 2}, n: 1, expected: []int{1}},
		{input: []int{1, 2}, n: 2, expected: []int{2}},
	}

	for i, tc := range tests {
		head := sliceToList(tc.input)
		got := listToSlice(removeNthFromEnd(head, tc.n))

		// normalize nil slices to empty
		if got == nil {
			got = []int{}
		}
		if tc.expected == nil {
			tc.expected = []int{}
		}

		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("Test %d failed: expected %v, got %v", i+1, tc.expected, got)
		}
	}
}
