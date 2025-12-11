package main

const mod int64 = 1_000_000_007

// countUnlockingPermutations returns the number of valid unlocking orders.
// 0 must be the unique minimum complexity; otherwise some node cannot be unlocked.
func countUnlockingPermutations(complexity []int) int {
	n := len(complexity)
	root := complexity[0]
	for i := 1; i < n; i++ {
		if complexity[i] <= root {
			return 0
		}
	}

	// Any order of the remaining (n-1) computers works.
	fact := int64(1)
	for i := 2; i <= n-1; i++ {
		fact = fact * int64(i) % mod
	}
	return int(fact)
}
