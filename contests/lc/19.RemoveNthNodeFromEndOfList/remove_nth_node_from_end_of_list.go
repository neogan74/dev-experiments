package _19_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy

	// продвигаем first на n+1 шагов
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// продвигаем оба указателя, пока first не достигнет конца
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// удаляем n-й с конца узел
	second.Next = second.Next.Next
	return dummy.Next
}
