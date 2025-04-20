package _563_CountTheNumberOfFairPairs

import "testing"

func TestCountFairPairs(t *testing.T) {
	tests := []struct {
		nums         []int
		lower, upper int
		expected     int64
	}{
		// Примеры из условия
		{[]int{0, 1, 7, 4, 4, 5}, 3, 6, 6},
		{[]int{1, 7, 9, 2, 5}, 11, 11, 1},

		// все элементы равны 2, пары, дающие сумму 4
		{[]int{2, 2, 2, 2}, 4, 4, 6}, // C(4,2)=6

		// негативные числа
		{[]int{-2, -1, 0, 1, 2}, -1, 1, 6},

		// пустой или слишком маленький массив
		{[]int{}, 0, 10, 0},
		{[]int{5}, 0, 10, 0},

		// широкие границы — всё
		{[]int{5, 1, 3}, -100, 100, 3}, // пары: (5,1),(5,3),(1,3)
	}

	for _, tt := range tests {
		got := countFairPairs(tt.nums, tt.lower, tt.upper)
		if got != tt.expected {
			t.Errorf(
				"countFairPairs(%v, %d, %d) = %d; want %d",
				tt.nums, tt.lower, tt.upper, got, tt.expected,
			)
		}
	}
}
