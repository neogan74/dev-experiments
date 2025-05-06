package _92_reverselinkedlistii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// Create a dummy node to handle edge case when left == 1 (i.e., reversing starts from head)
	dummy := &ListNode{Next: head}
	tempptr := dummy

	// Move tempptr to the node just before 'left' position
	for idx := 1; idx < left; idx++ {
		tempptr = tempptr.Next
	}

	// Reverse the sublist from tempptr.Next with length (right - left)
	reversedHead := reverse(tempptr.Next, right-left)

	// Connect the reversed sublist back to the original list
	tempptr.Next = reversedHead

	return dummy.Next
}

// Reverse function using a stack to reverse a sublist of given length
func reverse(head *ListNode, right int) *ListNode {
	var stack []*ListNode

	// No need to reverse if right == 0
	if right == 0 {
		return head
	}

	// Push (right + 1) nodes onto the stack
	for idx := 0; idx <= right; idx++ {
		stack = append(stack, head)
		head = head.Next
	}

	// Save the node following the reversed segment to reconnect later
	var nextLL *ListNode = head

	// Get length of the stack
	lenght := len(stack)

	// The new head of the reversed list will be the last node in the stack
	tempHead := stack[lenght-1]

	// Pop nodes and reverse the pointers
	for ; lenght > 0; lenght-- {
		if lenght == 1 {
			// The first node in original segment should point to nextLL
			stack[lenght-1].Next = nextLL
		} else {
			// Reverse the link: current node points to previous node in stack
			stack[lenght-1].Next = stack[lenght-2]
		}

		// Shrink the stack
		stack = stack[:lenght-1]
	}

	return tempHead
}
