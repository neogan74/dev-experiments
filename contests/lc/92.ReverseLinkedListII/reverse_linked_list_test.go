package _92_reverselinkedlistii

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

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		input    []int
		left     int
		right    int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, left: 2, right: 4, expected: []int{1, 4, 3, 2, 5}},
		{input: []int{5}, left: 1, right: 1, expected: []int{5}},
		{input: []int{1, 2, 3}, left: 1, right: 3, expected: []int{3, 2, 1}},
		{input: []int{1, 2, 3}, left: 2, right: 2, expected: []int{1, 2, 3}},
		{input: []int{1, 2, 3, 4}, left: 3, right: 4, expected: []int{1, 2, 4, 3}},
	}

	for i, tc := range tests {
		head := sliceToList(tc.input)
		got := listToSlice(reverseBetween(head, tc.left, tc.right))
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("Test %d failed: expected %v, got %v", i+1, tc.expected, got)
		}
	}
}
