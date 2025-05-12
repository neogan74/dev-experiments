package _61_rotate_list

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

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
		{[]int{1}, 0, []int{1}},
		{[]int{}, 1, []int{}},
		{[]int{1, 2}, 2, []int{1, 2}},
	}

	for i, tc := range tests {
		head := sliceToList(tc.input)
		got := listToSlice(rotateRight(head, tc.k))
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
