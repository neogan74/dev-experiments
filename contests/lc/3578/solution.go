package main

const mod int64 = 1_000_000_007

// countPartitions returns the number of valid partitions where every segment
// satisfies max - min <= k. Uses a sliding window with monotonic queues and
// prefix-sum DP for O(1) range-sum queries.
func countPartitions(nums []int, k int) int {
	n := len(nums)
	dp := make([]int64, n+1)   // dp[i]: ways to partition first i elements
	pref := make([]int64, n+1) // pref[i]: sum(dp[0..i])
	dp[0] = 1
	pref[0] = 1

	maxQ := make([]int, 0)
	minQ := make([]int, 0)
	left := 0

	for i, val := range nums {
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] < val {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)

		for len(minQ) > 0 && nums[minQ[len(minQ)-1]] > val {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, i)

		for int64(nums[maxQ[0]]-nums[minQ[0]]) > int64(k) {
			if maxQ[0] == left {
				maxQ = maxQ[1:]
			}
			if minQ[0] == left {
				minQ = minQ[1:]
			}
			left++
		}

		total := pref[i]
		if left > 0 {
			total = (total - pref[left-1] + mod) % mod
		} else {
			total %= mod
		}

		dp[i+1] = total
		pref[i+1] = (pref[i] + dp[i+1]) % mod
	}

	return int(dp[n] % mod)
}
