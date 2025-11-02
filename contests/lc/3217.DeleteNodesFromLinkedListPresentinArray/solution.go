package main

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// modifiedList removes nodes whose values appear in nums and returns the new head.
func modifiedList(nums []int, head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	blocked := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		blocked[v] = struct{}{}
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if _, banned := blocked[curr.Val]; banned {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return dummy.Next
}
