package mergeksortedlists

import (
	"reflect"
	"testing"
)

func makeList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{{}}, []int{}},
	}

	for _, tc := range tests {
		var lists []*ListNode
		for _, l := range tc.input {
			lists = append(lists, makeList(l))
		}
		got := listToSlice(mergeKLists(lists))
		if got == nil {
			got = []int{}
		}
		if tc.expected == nil {
			tc.expected = []int{}
		}
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("mergeKLists(%v) = %v; want %v", tc.input, got, tc.expected)
		}
	}
}
