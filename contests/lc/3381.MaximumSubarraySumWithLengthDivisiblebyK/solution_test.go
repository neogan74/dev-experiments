package _381_MaximumSubarraySumWithLengthDivisiblebyK

import (
	"math"
	"math/rand"
	"testing"
)

func bruteForce(nums []int, k int) int64 {
	best := int64(0)
	found := false
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if (j-i+1)%k == 0 {
				if !found || int64(sum) > best {
					best = int64(sum)
					found = true
				}
			}
		}
	}
	if !found {
		return math.MinInt64
	}
	return best
}

func TestExamples(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{1, 2}, 1, 3},
		{[]int{-1, -2, -3, -4, -5}, 4, -10},
		{[]int{-5, 1, 2, -3, 4}, 2, 4},
	}
	for _, tc := range tests {
		if got := maxSubarraySum(tc.nums, tc.k); got != tc.want {
			t.Fatalf("maxSubarraySum(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	if got := maxSubarraySum([]int{7}, 1); got != 7 {
		t.Fatalf("single element failed, got %d want 7", got)
	}

	nums := []int{-5, -4, -3, -2, -1}
	want := bruteForce(nums, 2)
	if got := maxSubarraySum(nums, 2); got != want {
		t.Fatalf("all negative failed, got %d want %d", got, want)
	}
}

func TestRandomSmallArrays(t *testing.T) {
	rng := rand.New(rand.NewSource(0))
	for i := 0; i < 100; i++ {
		n := rng.Intn(8) + 1
		k := rng.Intn(n) + 1
		nums := make([]int, n)
		for j := range nums {
			nums[j] = rng.Intn(11) - 5
		}
		want := bruteForce(nums, k)
		if got := maxSubarraySum(nums, k); got != want {
			t.Fatalf("case %d failed: nums=%v k=%d got %d want %d", i, nums, k, got, want)
		}
	}
}
