package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		result := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("twoSum(%v, %d) = %v; want %v",
				tt.nums, tt.target, result, tt.expected)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, -3, 8, -10, 12, 17, 23, 4, 16, 9, 5, 14, 18, 21, 3, 10, 6, 19, 13, 1, 11, 7}
	target := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
