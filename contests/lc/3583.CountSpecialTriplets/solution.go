package main

import "fmt"

const mod int64 = 1_000_000_007

// countSpecialTriplets returns the number of index triplets (i, j, k)
// such that i < j < k and nums[i] = nums[j]*2 and nums[k] = nums[j]*2.
// It uses prefix/suffix frequency maps to count matches in O(n).
func countSpecialTriplets(nums []int) int {
	suffix := make(map[int]int64, len(nums))
	for _, v := range nums {
		suffix[v]++
	}

	prefix := make(map[int]int64, len(nums))
	var ans int64

	for _, v := range nums {
		suffix[v]--
		target := v * 2

		left := prefix[target]
		right := suffix[target]
		ans = (ans + left*right) % mod

		prefix[v]++
	}

	return int(ans)
}

func specialTriplets(nums []int) (ans int) {
	const mod int = 1e9 + 7

	seen := [200_002]int{}
	seen2 := [200_002]int{}

	for _, num := range nums {
		ans += seen2[num]
		seen2[2*num] += seen[2*num]
		seen[num]++
	}

	ans = (ans%mod + mod) % mod
	return
}

func main() {
	samples := [][]int{
		{6, 3, 6},
		{0, 1, 0, 0},
		{8, 4, 2, 8, 4},
	}

	for _, nums := range samples {
		fmt.Printf("nums=%v -> %d\n", nums, countSpecialTriplets(nums))
		fmt.Printf("nums=%v -> %d\n", nums, specialTriplets(nums))
	}
}
