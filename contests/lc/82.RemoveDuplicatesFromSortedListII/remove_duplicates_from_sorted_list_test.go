package _82_remove_duplicates_from_sorted_list_ii

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

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 3, 4, 4, 5}, expected: []int{1, 2, 5}},
		{input: []int{1, 1, 1, 2, 3}, expected: []int{2, 3}},
	}
	for i, tc := range tests {
		head := sliceToList(tc.input)
		got := listToSlice(deleteDuplicates(head))
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("Test %d failed: expected %v, got %v", i+1, tc.expected, got)
		}
	}
}
