package _138_copy_list_with_random_pointers

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	oldToNew := make(map[*Node]*Node)

	curr := head
	for curr != nil {
		oldToNew[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		oldToNew[curr].Next = oldToNew[curr.Next]
		oldToNew[curr].Random = oldToNew[curr.Random]
		curr = curr.Next
	}

	return oldToNew[head]
}
