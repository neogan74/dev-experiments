package _338_CountTheNumberOfIdealArrays

// idealArrays возвращает число идеальных массивов длины n с элементами ≤ maxValue.
func idealArrays(n, maxValue int) int {
	mod := int64(1000000007)

	// binary exponentiation for FLT
	var exp func(int64, int64) int64
	exp = func(b int64, p int64) int64 {
		if p == 0 {
			return 1
		}
		if p == 1 {
			return b
		}
		half := exp(b, p/2)
		half *= half
		half %= mod
		if p%2 == 1 {
			half *= b
			half %= mod
		}
		return half
	}

	// precompute factorials up to 10000!
	fact := make([]int64, 10001)
	fact[0] = 1
	for i := 1; i < 10001; i++ {
		fact[i] = (fact[i-1] * int64(i)) % mod
	}
	// precompute modular inverses of factorials mod 10^9 + 7
	inv := make([]int64, 10001)
	inv[10000] = exp(fact[10000], mod-2)
	for i := 9999; i >= 0; i-- {
		inv[i] = (inv[i+1] * int64(i+1)) % mod
	}

	res := int64(0)

	dp := make([]int64, maxValue+1)
	// initialize to all 0 except for 1, so that 2 to maxValue initialize to 1 in first round of dp
	dp[1] = 1
	sum := int64(1)
	for jumps := 0; jumps < n && sum != 0; jumps++ {
		sum = int64(0)
		// generate number of sequences with the given number of jumps for each ending value
		dp2 := make([]int64, maxValue+1)
		for i := 1; i <= maxValue/2+1; i++ {
			// if we don't end up adding anything, skip
			if dp[i] == 0 {
				continue
			}
			for j := 2 * i; j <= maxValue; j += i {
				dp2[j] += dp[i]
				sum += dp[i]
			}
		}
		// for the first iteration, we need to make sure that dp2[1] = 1
		if jumps == 0 {
			sum++
			dp2[1] = 1
		}
		// binomial coefficient using precomputed values
		res += ((sum * fact[n-1]) % mod) * ((inv[n-1-jumps] * inv[jumps]) % mod) % mod
		res %= mod
		// update dp array to be used for next length of sequences
		dp = dp2
	}
	res %= mod
	return int(res)
}
