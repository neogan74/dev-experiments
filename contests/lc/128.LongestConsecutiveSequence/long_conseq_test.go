package _28_LongestConsecutiveSequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{}, 0},
		{[]int{1, 2, 0, 1}, 3},                        // [0,1,2]
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7}, // [-1,0,1,2,3,4,5,6]? actually 0..1.. then 3.. but the max is 8 consecutive from -1 to 6 => 8
		{[]int{5, 5, 5, 5}, 1},                        // дубликаты, только одна уникальная
	}

	for _, tt := range tests {
		got := longestConsecutive(tt.nums)
		if got != tt.expected {
			t.Errorf("longestConsecutive(%v) = %d; want %d", tt.nums, got, tt.expected)
		}
	}
}
