package _21_merge_two_sorted_lists

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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		{list1: []int{}, list2: []int{}, expected: []int{}},
		{list1: []int{}, list2: []int{0}, expected: []int{0}},
		{list1: []int{5, 6}, list2: []int{1, 2, 3}, expected: []int{1, 2, 3, 5, 6}},
	}

	for i, tc := range tests {
		l1 := sliceToList(tc.list1)
		l2 := sliceToList(tc.list2)
		got := listToSlice(mergeTwoLists(l1, l2))

		// Normalize nil slices to empty slices for deep equality
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
