package _19_ContainsDuplicateII

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 1}, 0, false},        // k=0, никакие два разных индекса не могут дать dist=0
		{[]int{}, 5, false},               // пустой массив
		{[]int{5, 5}, 1, true},            // сразу дубликат
		{[]int{5, 5}, 0, false},           // хотя значения равны, |i-j|=1>0
		{[]int{1, 2, 3, 4, 5}, 10, false}, // нет дубликатов
		{[]int{9, 9, 8, 9}, 2, true},      // пара (0,1) или (1,3)
	}

	for _, tt := range tests {
		got := containsNearbyDuplicate(tt.nums, tt.k)
		if got != tt.expected {
			t.Errorf("containsNearbyDuplicate(%v, %d) = %v; want %v",
				tt.nums, tt.k, got, tt.expected)
		}
	}
}
