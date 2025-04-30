package _141_linked_list_cycle

import "testing"

func Test_hasCycle(t *testing.T) {
	// Test 1: цикл есть
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2 // цикл
	if !hasCycle(n1) {
		t.Errorf("TestHasCycle failed: expected true, got false")
	}
	// Test 2: цикл отсутствует
	a1 := &ListNode{Val: 1}
	a2 := &ListNode{Val: 2}
	a1.Next = a2
	if hasCycle(a1) {
		t.Errorf("TestHasCycle failed: expected false, got true")
	}
	// Test 3: один узел, без цикла
	solo := &ListNode{Val: 1}
	if hasCycle(solo) {
		t.Errorf("TestHasCycle failed: expected false, got true")
	}
	// Test 4: один узел, с циклом на самого себя
	self := &ListNode{Val: 1}
	self.Next = self
	if !hasCycle(self) {
		t.Errorf("TestHasCycle failed: expected true, got false")
	}
	// Test 5: пустой список
	if hasCycle(nil) {
		t.Errorf("TestHasCycle failed: expected false, got true")
	}
}
