package _3068_Find_the_Maximum_Sum_of_Node_Values

import "testing"

func TestMaximumValueSum(t *testing.T) {
	nums := []int{24, 78, 1, 97, 44}
	k := 6
	edges := [][]int{{0, 2}, {1, 2}, {4, 2}, {3, 4}}
	var expected int64 = 260

	got := maximumValueSum(nums, k, edges)
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
