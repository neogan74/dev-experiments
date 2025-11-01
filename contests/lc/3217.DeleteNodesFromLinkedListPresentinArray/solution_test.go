package main

import "testing"

func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	curr := head
	for _, v := range values[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func TestModifiedList(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		list     []int
		expected []int
	}{
		{
			name:     "remove scattered values",
			nums:     []int{1, 4, 3},
			list:     []int{2, 1, 4, 2, 3, 7},
			expected: []int{2, 2, 7},
		},
		{
			name:     "remove head nodes",
			nums:     []int{5, 6},
			list:     []int{5, 6, 7, 8},
			expected: []int{7, 8},
		},
		{
			name:     "remove nothing",
			nums:     []int{9, 10},
			list:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "remove all nodes",
			nums:     []int{1, 2, 3},
			list:     []int{1, 2, 3, 2},
			expected: nil,
		},
		{
			name:     "duplicate values in nums",
			nums:     []int{2, 2, 3},
			list:     []int{1, 2, 3, 4},
			expected: []int{1, 4},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.list)
			got := modifiedList(tc.nums, head)
			if actual := listToSlice(got); !slicesEqual(actual, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkModifiedList(b *testing.B) {
	nums := make([]int, 500)
	for i := range nums {
		nums[i] = i
	}

	listVals := make([]int, 10000)
	for i := range listVals {
		listVals[i] = i % 750
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := buildList(listVals)
		modifiedList(nums, head)
	}
}
