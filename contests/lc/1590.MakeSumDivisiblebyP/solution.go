package main

// minSubarray finds the shortest subarray to remove so the remaining sum is divisible by p.
func minSubarray(nums []int, p int) int {
	// total modulo p; we only need the remainder, not the full sum.
	total := 0
	for _, v := range nums {
		total = (total + v%p) % p
	}
	target := total % p
	if target == 0 {
		return 0
	}

	n := len(nums)
	best := n + 1
	prefix := 0
	pos := map[int]int{0: -1} // remainder -> earliest index

	for i, v := range nums {
		prefix = (prefix + v%p) % p
		need := (prefix - target + p) % p
		if idx, ok := pos[need]; ok && i-idx < best {
			best = i - idx
		}
		pos[prefix] = i
	}

	if best < n {
		return best
	}
	return -1
}

func minSubarray2(nums []int, p int) int {
	lastIdx := make(map[int]int, len(nums))
	lastIdx[0] = -1

	sum := 0
	pref := 0

	for _, v := range nums {
		sum = (sum + v) % p
	}

	ans := len(nums)
	for i := 0; i <= len(nums); i++ {
		suf := (sum - pref + p) % p

		need := (p - suf) % p
		if last, ok := lastIdx[need]; ok {
			ans = min(ans, i-last-1)
		}

		if i < len(nums) {
			pref = (pref + nums[i]) % p
			lastIdx[pref] = i
		}
	}

	if ans == len(nums) {
		return -1
	}

	return ans
}
