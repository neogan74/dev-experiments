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

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	copyHead := head.Next
	copyCur := copyHead
	for cur != nil {
		cur.Next = cur.Next.Next
		if copyCur.Next != nil {
			copyCur.Next = copyCur.Next.Next
		}
		cur = cur.Next
		copyCur = copyCur.Next
	}

	return copyHead
}
